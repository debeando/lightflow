package mysql_test

import (
	"testing"

	"github.com/debeando/lightflow/flow/mysql"
)

func TestValidPath(t *testing.T) {
	m := mysql.MySQL{
		Path: "/tmp/test.c",
	}

	t.Errorf("%s", m.ValidPath())
}
