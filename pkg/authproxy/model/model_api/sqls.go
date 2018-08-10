package model_api

var (
	sqlSelectUser string = "SELECT u.name, u.password, r.name as role_name, p.name as project_name, pr.name as project_role_name FROM users as u " +
		"INNER JOIN user_roles as ur ON u.id = ur.user_id " +
		"INNER JOIN roles as r ON ur.role_id = r.id " +
		"INNER JOIN projects as p ON r.project_id = p.id " +
		"INNER JOIN project_roles as pr ON p.project_role_id = pr.id "
)
