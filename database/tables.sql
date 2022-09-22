CREATE TABLE users (
	id 	 SERIAL PRIMARY KEY,
	name VARCHAR(255) UNIQUE
);

CREATE TABLE trips (
	id 			 SERIAL PRIMARY KEY,
	time_stamp 	 TIMESTAMPTZ,
	user_id		 INTEGER REFERENCES users(id),
	distance 	 INTEGER,
	gallons_used INTEGER,
	cost_of_fuel FLOAT,
	start_dest 	 VARCHAR(255),
	end_dest 	 VARCHAR(255),

	unique(user_id, time_stamp)
);