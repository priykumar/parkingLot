package repository

import (
	"fmt"
	"strings"
)

func (p *ParkingLotRepository) CreateTables() error {
	query := `
		CREATE TABLE IF NOT EXISTS parking_lot (
			id INTEGER NOT NULL AUTO_INCREMENT,
			address TEXT,
			PRIMARY KEY (id)
		);`

	_, err := p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table parking_lot, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS electric_charger (
		id INTEGER NOT NULL AUTO_INCREMENT,
		consumption INTEGER DEFAULT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table electric_charger, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS operator (
		id INTEGER NOT NULL AUTO_INCREMENT,
		name TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table operator, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS gate_status (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table gate_status, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS gate_type (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table gate_type, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS invoice_paid_status (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table invoice_paid_status, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS parking_spot_status (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table parking_spot_status, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS payment_mode (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table payment_mode, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS payment_status (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table payment_status, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS vehicle_type (
		id INTEGER NOT NULL AUTO_INCREMENT,
		value TEXT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table vehicle_type, Error: ", err)
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS parking_floor (
			id INTEGER NOT NULL AUTO_INCREMENT,
			floor_number INTEGER NOT NULL,
			parking_lot_id INTEGER NOT NULL,
			PRIMARY KEY (id),
			KEY parking_floor_FK (parking_lot_id),
			CONSTRAINT parking_floor_FK FOREIGN KEY (parking_lot_id) REFERENCES parking_lot (id)
		);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table parking_floor, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS gate (
		id INTEGER NOT NULL AUTO_INCREMENT,
		gate_number INTEGER NOT NULL,
		gate_type_id INTEGER NOT NULL,
		parking_lot_id INTEGER NOT NULL,
		gate_status_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY gate2_FK (gate_type_id),
		KEY gate_FK (parking_lot_id),
		KEY gate_FK_1 (gate_status_id),
		CONSTRAINT gate2_FK FOREIGN KEY (gate_type_id) REFERENCES gate_type (id),
		CONSTRAINT gate_FK FOREIGN KEY (parking_lot_id) REFERENCES parking_lot (id),
		CONSTRAINT gate_FK_1 FOREIGN KEY (gate_status_id) REFERENCES gate_status (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table gate, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS parking_spot (
		id INTEGER NOT NULL AUTO_INCREMENT,
		spot_number INTEGER NOT NULL,
		parking_floor_id INTEGER NOT NULL,
		vehicle_type_id INTEGER NOT NULL,
		parking_spot_status_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY parking_spot_FK (parking_floor_id),
		KEY parking_spot_FK_2 (parking_spot_status_id),
		KEY parking_spot_FK_1 (vehicle_type_id),
		CONSTRAINT parking_spot_FK FOREIGN KEY (parking_floor_id) REFERENCES parking_floor (id),
		CONSTRAINT parking_spot_FK_1 FOREIGN KEY (vehicle_type_id) REFERENCES vehicle_type (id),
		CONSTRAINT parking_spot_FK_2 FOREIGN KEY (parking_spot_status_id) REFERENCES parking_spot_status (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table parking_spot, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS electric_parking_spot (
		id INTEGER NOT NULL AUTO_INCREMENT,
		parking_spot_id INTEGER NOT NULL,
		electric_charger_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY electric_parking_spot_FK (parking_spot_id),
		KEY electric_parking_spot_FK_1 (electric_charger_id),
		CONSTRAINT electric_parking_spot_FK FOREIGN KEY (parking_spot_id) REFERENCES parking_spot (id),
		CONSTRAINT electric_parking_spot_FK_1 FOREIGN KEY (electric_charger_id) REFERENCES electric_charger (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table electric_parking_spot, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS vehicle (
		id INTEGER NOT NULL AUTO_INCREMENT,
		vehicle_number INTEGER NOT NULL,
		vehicle_type_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY vehicle_FK (vehicle_type_id),
		CONSTRAINT vehicle_FK FOREIGN KEY (vehicle_type_id) REFERENCES vehicle_type (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table vehicle, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS ticket (
		id INTEGER NOT NULL AUTO_INCREMENT,
		entry_time datetime NOT NULL,
		entry_gate_id INTEGER NOT NULL,
		operator_id INTEGER NOT NULL,
		vehicle_id INTEGER NOT NULL,
		parking_spot_id INTEGER NOT NULL,
		parking_lot_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY ticket_FK (entry_gate_id),
		KEY ticket_FK_1 (operator_id),
		KEY ticket_FK_2 (vehicle_id),
		KEY ticket_FK_3 (parking_spot_id),
		KEY ticket_FK_4 (parking_lot_id),
		CONSTRAINT ticket_FK FOREIGN KEY (entry_gate_id) REFERENCES gate (id),
		CONSTRAINT ticket_FK_1 FOREIGN KEY (operator_id) REFERENCES operator (id),
		CONSTRAINT ticket_FK_2 FOREIGN KEY (vehicle_id) REFERENCES vehicle (id),
		CONSTRAINT ticket_FK_3 FOREIGN KEY (parking_spot_id) REFERENCES parking_spot (id),
		CONSTRAINT ticket_FK_4 FOREIGN KEY (parking_lot_id) REFERENCES parking_lot (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table ticket, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS invoice (
		id INTEGER NOT NULL AUTO_INCREMENT,
		amount INTEGER NOT NULL,
		exit_time datetime NOT NULL,
		ticket_id INTEGER NOT NULL,
		operator_id INTEGER NOT NULL,
		invoice_paid_status_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY invoice_FK (ticket_id),
		KEY invoice_FK_1 (invoice_paid_status_id),
		KEY invoice_FK_2 (operator_id),
		CONSTRAINT invoice_FK FOREIGN KEY (ticket_id) REFERENCES ticket (id),
		CONSTRAINT invoice_FK_1 FOREIGN KEY (invoice_paid_status_id) REFERENCES invoice_paid_status (id),
		CONSTRAINT invoice_FK_2 FOREIGN KEY (operator_id) REFERENCES operator (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table invoice, Error: ", err)
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS payment (
		id INTEGER NOT NULL AUTO_INCREMENT,
		referance_number INTEGER NOT NULL,
		pay_time datetime NOT NULL,
		amount INTEGER NOT NULL,
		invoice_id INTEGER NOT NULL,
		payment_mode_id INTEGER NOT NULL,
		payment_status_id INTEGER NOT NULL,
		PRIMARY KEY (id),
		KEY payment_FK (invoice_id),
		KEY payment_FK_1 (payment_mode_id),
		KEY payment_FK_2 (payment_status_id),
		CONSTRAINT payment_FK FOREIGN KEY (invoice_id) REFERENCES invoice (id),
		CONSTRAINT payment_FK_1 FOREIGN KEY (payment_mode_id) REFERENCES payment_mode (id),
		CONSTRAINT payment_FK_2 FOREIGN KEY (payment_status_id) REFERENCES payment_status (id)
	);`
	_, err = p.DbClient.Exec(query)
	if err != nil {
		fmt.Println("Failed creating table payment, Error: ", err)
		return err
	}

	fmt.Println("Tables are successfully created")

	return nil
}

func (p *ParkingLotRepository) InsertValuesIntoEnums() error {

	// gate_status
	enumMap := map[string][]string{
		"gate_status":         {"OPEN", "CLOSED"},
		"gate_type":           {"ENTRY", "EXIT"},
		"invoice_paid_status": {"PAID", "UNPAID"},
		"parking_spot_status": {"AVAILABLE", "UNAVAILABLE"},
		"payment_mode":        {"CASH", "DEBIT_CARD", "CREDIT_CARD", "NET_BANKING", "UPI"},
		"payment_status":      {"SUCCESS", "PENDING", "FAILURE"},
		"vehicle_type":        {"SMALL", "MEDIUM", "LARGE", "ELECTRIC"},
	}

	for k, v := range enumMap {
		cmd := fmt.Sprintf("INSERT INTO %s (value) VALUES (?)", k)
		for _, vv := range v {
			_, err := p.DbClient.Exec(cmd, vv)
			if err != nil {
				fmt.Printf("Error inserting %s in %s. Error: ", vv, k)
				return err
			}
		}
	}
	fmt.Println("All the enum tables are populated")
	return nil
}

func (p *ParkingLotRepository) DeleteTables() error {
	deleteOrder := []string{"electric_parking_spot",
		"payment",
		"invoice",
		"ticket",
		"parking_spot",
		"parking_floor",
		"gate",
		"parking_lot",
		"electric_charger",
		"gate_status",
		"gate_type",
		"invoice_paid_status",
		"operator",
		"parking_spot_status",
		"payment_mode",
		"payment_status",
		"vehicle",
		"vehicle_type"}

	deletedTableList := ""
	for i := range deleteOrder {
		cmd := fmt.Sprintf("DROP TABLE %s;", deleteOrder[i])
		_, err := p.DbClient.Exec(cmd)
		if err != nil {
			if strings.Contains(err.Error(), "Unknown table") {
				continue
			}
			return err
		} else {
			deletedTableList = deletedTableList + deleteOrder[i] + " "
		}

	}
	if deletedTableList != "" {
		fmt.Printf("Tables [ %s] are deleted from DB\n", deletedTableList)
	}
	return nil
}
