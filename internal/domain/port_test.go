package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPort(t *testing.T) {
	t.Parallel()

	portId := "port id"
	portName := "port name"
	portCode := "port code"
	portCity := "port city"
	portCountry := "port country"

	testCases := []struct {
		name    string
		port    func() (*Port, error)
		isValid bool
	}{
		{
			name: "valid",
			port: func() (*Port, error) {
				return NewPort(portId, portName, portCode, portCity, portCountry,
					nil, nil, nil, "", "", nil)
			},
			isValid: true,
		},
		{
			name: "missing port Id",
			port: func() (*Port, error) {
				return NewPort("", portName, portCode, portCity, portCountry,
					nil, nil, nil, "", "", nil)
			},
			isValid: false,
		},
		{
			name: "missing port name",
			port: func() (*Port, error) {
				return NewPort(portId, "", portCode, portCity, portCountry,
					nil, nil, nil, "", "", nil)
			},
			isValid: false,
		},
		{
			name: "missing port city",
			port: func() (*Port, error) {
				return NewPort(portId, portName, portCode, "", portCountry,
					nil, nil, nil, "", "", nil)
			},
			isValid: false,
		},
		{
			name: "missing port country",
			port: func() (*Port, error) {
				return NewPort(portId, portName, portCode, portCity, "",
					nil, nil, nil, "", "", nil)
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			port, err := tc.port()
			if tc.isValid {
				assert.NoError(t, err)

				assert.Equal(t, portId, port.Id())
				assert.Equal(t, portCode, port.Code())
				assert.Equal(t, portName, port.Name())
				assert.Equal(t, portCity, port.City())
				assert.Equal(t, portCountry, port.Country())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
