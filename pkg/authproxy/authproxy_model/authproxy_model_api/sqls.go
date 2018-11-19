package authproxy_model_api

var (
	sqlSelectUser string = "SELECT u.name, r.id as role_id, r.name as role_name, p.name as project_name, pr.id as project_role_id, pr.name as project_role_name, s.id as service_id, s.name as service_name, s.scope as service_scope " +
		"FROM users as u " +
		"INNER JOIN user_roles as ur ON u.id = ur.user_id " +
		"INNER JOIN roles as r ON ur.role_id = r.id " +
		"INNER JOIN projects as p ON r.project_id = p.id " +
		"INNER JOIN project_roles as pr ON p.project_role_id = pr.id " +
		"INNER JOIN project_role_services as prs ON pr.id = prs.project_role_id " +
		"INNER JOIN services as s ON prs.service_id = s.id "
)