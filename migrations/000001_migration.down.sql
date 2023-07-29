-- +migrate Down
DROP TABLE IF NOT EXISTS data;
DROP TABLE IF NOT EXISTS Car_Service_quarter;
DROP TABLE IF NOT EXISTS Car_Service_month;
DROP TABLE IF NOT EXISTS Tachograph_Services_Report;
DROP TABLE IF NOT EXISTS AutoElectrician_Services_Report;
DROP TABLE IF NOT EXISTS heating_services;
DROP TABLE IF NOT EXISTS water_check;