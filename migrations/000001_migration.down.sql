-- +migrate Down
DROP TABLE IF EXISTS data;
DROP TABLE IF EXISTS Car_Service_quarter;
DROP TABLE IF EXISTS Car_Service_month;
DROP TABLE IF EXISTS Tachograph_Services_Report;
DROP TABLE IF EXISTS AutoElectrician_Services_Report;
DROP TABLE IF EXISTS heating_services;
DROP TABLE IF EXISTS water_check;