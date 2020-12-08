terraform {
  required_providers {
    dbor = {
      source = "nvmathure/sqldb/dbor"
      version = "0.0.1"
    }
  }
}

provider "dbor" {

    password = "password"
    datasource = "datasource"
}

resource "dbor" "table" {
  name = Employeess
  columns = [
    {
      name = EmployeeId
      type = raw
      lenght = 16
    },
    {
      name = FirstName
      type = varchar2
      lenght = 100
    }
  ]
}