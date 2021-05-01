// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rodrigo-brito/ninjabot/pkg/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ExchangeID holds the value of the "exchange_id" field.
	ExchangeID int64 `json:"exchange_id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Symbol holds the value of the "symbol" field.
	Symbol string `json:"symbol,omitempty"`
	// Side holds the value of the "side" field.
	Side string `json:"side,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity float64 `json:"quantity,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int64 `json:"group_id,omitempty"`
	// Stop holds the value of the "stop" field.
	Stop float64 `json:"stop,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldPrice, order.FieldQuantity, order.FieldStop:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldExchangeID, order.FieldGroupID:
			values[i] = new(sql.NullInt64)
		case order.FieldSymbol, order.FieldSide, order.FieldType, order.FieldStatus:
			values[i] = new(sql.NullString)
		case order.FieldDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int64(value.Int64)
		case order.FieldExchangeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exchange_id", values[i])
			} else if value.Valid {
				o.ExchangeID = value.Int64
			}
		case order.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				o.Date = value.Time
			}
		case order.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				o.Symbol = value.String
			}
		case order.FieldSide:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field side", values[i])
			} else if value.Valid {
				o.Side = value.String
			}
		case order.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = value.String
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = value.String
			}
		case order.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				o.Price = value.Float64
			}
		case order.FieldQuantity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				o.Quantity = value.Float64
			}
		case order.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				o.GroupID = value.Int64
			}
		case order.FieldStop:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stop", values[i])
			} else if value.Valid {
				o.Stop = value.Float64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", exchange_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ExchangeID))
	builder.WriteString(", date=")
	builder.WriteString(o.Date.Format(time.ANSIC))
	builder.WriteString(", symbol=")
	builder.WriteString(o.Symbol)
	builder.WriteString(", side=")
	builder.WriteString(o.Side)
	builder.WriteString(", type=")
	builder.WriteString(o.Type)
	builder.WriteString(", status=")
	builder.WriteString(o.Status)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", o.Price))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", o.Quantity))
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", o.GroupID))
	builder.WriteString(", stop=")
	builder.WriteString(fmt.Sprintf("%v", o.Stop))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
