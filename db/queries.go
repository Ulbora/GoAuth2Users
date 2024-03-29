//Package db ...
package db

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

const (
	connectTest = "select count(*) from user "

	insertRole  = "INSERT INTO role (role) values(?) "
	getRole     = "SELECT * FROM role WHERE id = ?"
	getRoleList = "SELECT * FROM role"
	deleteRole  = "DELETE FROM role WHERE id = ?"

	insertUser = "INSERT INTO user (username, password, enabled, date_entered, email_address, first_name, " +
		"last_name, role_id, client_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?) "

	updateUser = "UPDATE user SET password = ?, enabled = ?, email_address = ?, first_name = ?, " +
		"last_name = ?, role_id = ? " +
		"WHERE username = ? and client_id = ? "

	getUser = "SELECT * FROM user WHERE username = ? and client_id = ? "

	getUserList = "SELECT username, password, enabled, date_entered, email_address, first_name, " +
		"last_name, role_id, client_id " +
		"FROM user " +
		"order by client_id "

	searchUserList = "SELECT username, password, enabled, date_entered, email_address, first_name, " +
		"last_name, role_id, client_id " +
		"FROM user " +
		"where client_id = ? " +
		"order by username "

	deleteUser = "DELETE FROM user WHERE username = ? and client_id = ? "
)
