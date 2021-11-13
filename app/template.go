package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

var (
	sql   = `insert into %s values(%s);`
	table = `
create table %s
(
	%s
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
`
)

func Xlsx2MySQL(filename, sheet string) (string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return "", errors.New("open file error")
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return "", errors.New("it don't find sheet name")
	}
	fileName := strings.Split(filename, ".")[0]
	res, err := GenerateTableAndSQL(fileName, rows)
	if err != nil {
		return "", errors.New("generate sql fail")
	}
	err = OutputSQLFile("./"+fileName+".sql", res)
	return res, err
}

func GenerateTableAndSQL(tableName string, rows [][]string) (s string, err error) {
	sqlStr := ""
	resultSQL := ""
	if len(rows) < 2 {
		return sqlStr, errors.New("rows at least 2")
	}
	tableStr := ""
	temp := ""

	for i, row := range rows {
		for j, colCell := range row {
			// 第一第二行处理属性以及属性的类型, 第三行处理注释
			if i == 0 {
				strTemp := (warpperString(colCell) + " %s comment '%s',\n  ")
				temp = temp + fmt.Sprintf(strTemp, rows[1][j], rows[2][j])
			}

			// 从第四行开始处理插入sql的语句，将他们拼接到一起
			// if i == 3 {
			// 	if j == 0 {
			// 		sqlStr += ""
			// 	} else {
			// 		sqlStr += ","
			// 	}
			// }

			if i >= 3 {
				if j == 0 {
					sqlStr += ""
				} else {
					sqlStr += ","
				}
				attr := rows[1][j]
				if strings.Contains(attr, "int") {
					sqlStr = sqlStr + colCell
				} else {
					sqlStr = sqlStr + "'" + colCell + "'"
				}
				if len(rows[1]) == (j + 1) {
					resultSQL = resultSQL + fmt.Sprintf(sql, warpperString(tableName), sqlStr) + "\n"
					sqlStr = ""
				}
			}

		}

	}

	tableStr = fmt.Sprintf(table, warpperString(tableName), strings.TrimRight(temp, ",\n "))

	res := tableStr + resultSQL
	return res, nil
}

func OutputSQLFile(outputFilePath, text string) error {
	var err error
	f, err := os.Create(outputFilePath)
	if err != nil {
		return errors.New("create file error")
	}
	_, err = f.WriteString(text)
	return err
}

func warpperString(str string) string {
	return "`" + str + "`"
}
