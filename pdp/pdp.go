package pdp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/loader"
	"github.com/open-policy-agent/opa/rego"

	"github.com/kg0r0/authzen-interop-pdp-example/pip"
)

const (
	path = "./policy/todoapp.rego"
)

// Ref: https://openid.github.io/authzen/authorization-api-1_1_01#name-information-model
type AccessEvaluationAPIRequest struct {
	Subject struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Properties struct {
			Department string `json:"department,omitempty"`
			IPAddress  string `json:"ip_address,omitempty"`
			DeviceID   string `json:"device_id,omitempty"`
		} `json:"properties,omitempty"`
	} `json:"subject"`
	Resource struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Properties struct {
			LibraryRecord struct {
				Title string `json:"title,omitempty"`
				Isbn  string `json:"isbn,omitempty"`
			} `json:"library_record,omitempty"`
			OwnerID string `json:"ownerID,omitempty"`
		} `json:"properties,omitempty"`
	} `json:"resource"`
	Action struct {
		Name       string `json:"name"`
		Properties struct {
			Method string `json:"method"`
		} `json:"properties"`
	} `json:"action"`
	Context struct {
		Time string `json:"time,omitempty"`
	} `json:"context,omitempty"`
}

// Ref: https://openid.net/specs/authorization-api-1_0-01.html#section-6.2
type AccessEvaluationAPIResponse struct {
	Decision bool `json:"decision"`
	Context  struct {
		ID          string `json:"id,omitempty"`
		ReasonAdmin struct {
			En string `json:"en,omitempty"`
		} `json:"reason_admin,omitempty"`
		ReasonUser struct {
			En403 string `json:"en-403,omitempty"`
			Es403 string `json:"es-403,omitempty"`
		} `json:"reason_user,omitempty"`
	} `json:"context,omitempty"`
}

func Evaluation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method != http.MethodPost {
		slog.Info("Method not allowed", "method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Failed to read request body", "error", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var aer AccessEvaluationAPIRequest
	err = json.Unmarshal(body, &aer)
	if err != nil {
		slog.Error("Failed to unmarshal request body", "error", err)
		jsonResponse(w, AccessEvaluationAPIResponse{Decision: false}, http.StatusBadRequest)
		return
	}
	slog.Info("The Access Evaluation API Request", "body", string(body))

	attr := pip.UserAttributes[aer.Subject.ID]
	slog.Info("The user attributes", "attributes", attr)

	// Ref: https://authzen-interop.net/docs/scenarios/todo-1.1/#subjects
	regoFile, err := loader.RegoWithOpts(path, ast.ParserOptions{})
	if err != nil {
		slog.Error("Failed to load rego file", "error", err)
		jsonResponse(w, AccessEvaluationAPIResponse{Decision: false}, http.StatusInternalServerError)
		return
	}

	compiler, err := ast.CompileModules(map[string]string{
		"example.rego": regoFile.Parsed.String(),
	})
	if err != nil {
		slog.Error("Failed to compile rego file", "error", err)
		jsonResponse(w, AccessEvaluationAPIResponse{Decision: false}, http.StatusInternalServerError)
		return
	}

	rego := rego.New(
		rego.Query("data.policy.decision"),
		rego.Compiler(compiler),
		rego.Input(map[string]interface{}{
			"subject":  aer.Subject,
			"action":   aer.Action,
			"resource": aer.Resource,
			"email":    attr.Email,
		}),
	)

	rs, err := rego.Eval(ctx)
	if err != nil {
		slog.Error("Failed to evaluate rego file", "error", err)
		jsonResponse(w, AccessEvaluationAPIResponse{Decision: false}, http.StatusInternalServerError)
		return
	}
	decision := rs.Allowed()
	slog.Info("The Access Evaluation API Response", "decision", decision)
	jsonResponse(w, AccessEvaluationAPIResponse{Decision: decision}, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
