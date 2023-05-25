package util

func URDUtil(sql string, parameterList ...interface{}) (int, error) {

	stmt, err := GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	exec, err := stmt.Exec(parameterList...)
	if err != nil {
		return 0, err
	}
	i, _ := exec.RowsAffected()
	return int(i), nil

}

func URDUtil2(sql string, parameterList ...any) (int, error) {

	stmt, err := GetDbSongTest().Prepare(sql)
	defer stmt.Close()
	exec, err := stmt.Exec(parameterList...)
	if err != nil {
		return 0, err
	}
	i, _ := exec.RowsAffected()
	return int(i), nil

}
func MalysiaURDUtil(sql string, parameterList ...any) (int, error) {

	stmt, err := GetDbMalysia().Prepare(sql)
	defer stmt.Close()
	exec, err := stmt.Exec(parameterList...)
	if err != nil {
		return 0, err
	}
	i, _ := exec.RowsAffected()
	return int(i), nil

}
func MalysiaURDUtilReturnSongId(sql string, parameterList ...any) (int, error) {
	stmt, err := GetDbMalysia().Prepare(sql)
	defer stmt.Close()
	exec, err := stmt.Exec(parameterList...)
	if err != nil {
		return 0, err
	}
	i, _ := exec.LastInsertId()
	return int(i), nil

}
