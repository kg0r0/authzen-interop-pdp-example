package policy

import rego.v1

default decision := false

decision if {
	roles[input.subject.id][_] == "viewer"
	actions.viewer[_] == input.action.name
}

decision if {
    roles[input.subject.id][_] == "editor"
    actions.editor[_] == input.action.name
    recource_ckeck
}

resource_ckeck if {
    not input.resource.properties
}

recource_ckeck if {
    input.resource.properties
    not input.resource.properties.ownerID
}

recource_ckeck if {
    input.resource.properties.ownerID == input.email
}

decision if {
	roles[input.subject.id][_] == "admin"
	actions.admin[_] == input.action.name
}

decision if {
	roles[input.subject.id][_] == "evil_genius"
	actions.evil_genius[_] == input.action.name
}

roles := {
	# Rick Sanchez
	"CiRmZDA2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs": ["admin", "evil_genius"],
	# Beth Smith
	"CiRmZDM2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs": ["viewer"],
	# Morty Smith
	"CiRmZDE2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs": ["editor"],
	# Summer Smith
	"CiRmZDI2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs": ["editor"],
	# Jerry Smith
	"CiRmZDQ2MTRkMy1jMzlhLTQ3ODEtYjdiZC04Yjk2ZjVhNTEwMGQSBWxvY2Fs": ["viewer"],
}

actions := {
	"viewer": ["can_read_todos", "can_read_user"],
	"editor": ["can_read_todos", "can_read_user", "can_create_todo", "can_delete_todo", "can_update_todo"],
	"admin": ["can_read_todos", "can_read_user", "can_create_todo", "can_delete_todo", "can_update_todo"],
	"evil_genius": ["can_read_todos", "can_read_user", "can_create_todo", "can_update_todo"],
}
