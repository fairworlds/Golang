package fixapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	employee, err := ReadJSON("data.json")
	assert.NoError(t, err, "нет ошибок в JSON")
	assert.NotNil(t, employee, "Employees not nil")

	Emp := Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3}
	assert.Equal(t, Emp, employee[0])
	//	}
}
