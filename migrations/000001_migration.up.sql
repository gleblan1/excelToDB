-- +migrate Up
CREATE TABLE IF NOT EXISTS data (
    id INT PRIMARY KEY AUTO_INCREMENT,
    column1 VARCHAR(255) NOT NULL,
    column2 VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS water_check (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE NOT NULL,
    employee VARCHAR(255) NOT NULL,
    work_cost DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    total_quantity DECIMAL(10, 2) NOT NULL,
    total_cost DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS heating_services (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Date DATE NOT NULL,
    Client VARCHAR(255) NOT NULL,
    Employee VARCHAR(255) NOT NULL,
    Manager VARCHAR(255) NOT NULL,
    Deal_Source VARCHAR(255),
    Payment BOOLEAN NOT NULL ,
    Materials_Cost DECIMAL(10, 2),
    Installation DECIMAL(10, 2),
    Diagnostics DECIMAL(10, 2),
    Repairs DECIMAL(10, 2),
    Total_Amount DECIMAL(10, 2) NOT NULL,
    Total_Materials_Cost DECIMAL(10, 2),
    Total_Installation DECIMAL(10,2),
    Total_Diagnostics DECIMAL(10, 2),
    Total_Repairs DECIMAL(10, 2),
    Total DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS AutoElectrician_Services_Report (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Date DATE NOT NULL,
    Client VARCHAR(255) NOT NULL,
    Employee VARCHAR(255) NOT NULL,
    Deal_Source VARCHAR(255),
    Payment BOOLEAN NOT NULL,
    Materials_Cost DECIMAL(10, 2),
    AutoElectrician DECIMAL(10, 2),
    Starter DECIMAL(10, 2),
    Generator DECIMAL(10, 2),
    Headlight_Adjustment DECIMAL(10, 2),
    Injectors DECIMAL(10, 2),
    Total_Amount DECIMAL(10, 2) NOT NULL,
    Total_AutoElectrician DECIMAL(10, 2),
    Total_Starter DECIMAL(10, 2),
    Total_Generator DECIMAL(10, 2),
    Total_Headlight_Adjustment DECIMAL(10, 2),
    Total_Injectors DECIMAL(10, 2),
    Total_Materials_Cost DECIMAL(10,2),
    Total DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS Tachograph_Services_Report (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Date DATE NOT NULL,
    Payer VARCHAR(255) NOT NULL,
    Employee VARCHAR(255) NOT NULL,
    Manager VARCHAR(255) NOT NULL,
    Deal_Source VARCHAR(255),
    Payment BOOLEAN NOT NULL,
    Materials_Cost DECIMAL(10, 2),
    Installation DECIMAL(10, 2),
    Verification DECIMAL(10, 2),
    Repairs DECIMAL(10, 2),
    Total_Amount DECIMAL(10, 2) NOT NULL,
    Total_Installation DECIMAL(10, 2),
    Total_Verification DECIMAL(10, 2),
    Total_Repairs DECIMAL(10, 2),
    Total_Materials_Cost DECIMAL(10,2),
    Total DECIMAL(10, 2) NOT NULL
);



CREATE TABLE IF NOT EXISTS Car_Service_month (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Registration_Number INT NOT NULL,
    Date DATE NOT NULL,
    Car_Owner VARCHAR(255) NOT NULL,
    Car_Make VARCHAR(255) NOT NULL,
    License_Plate VARCHAR(20) NOT NULL,
    Category VARCHAR(50) NOT NULL,
    Quantity INT NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    Total_Quantity INT NOT NULL,
    Total_Amount DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS Car_Service_quarter (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Registration_Number INT NOT NULL,
    Date DATE NOT NULL,
    Work_Type VARCHAR(255) NOT NULL,
    Car_Owner VARCHAR(255) NOT NULL,
    Owner_Address VARCHAR(255) NOT NULL,
    Car_Make VARCHAR(255) NOT NULL,
    License_Plate VARCHAR(20) NOT NULL,
    Category VARCHAR(50) NOT NULL,
    Tachograph_Number VARCHAR(50) NOT NULL,
    Tachograph_Type VARCHAR(50) NOT NULL,
    Owned VARCHAR(255) NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    Total_Amount DECIMAL(10, 2) NOT NULL
);





