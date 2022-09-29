-- trips made by users
CREATE TABLE trips (
	id 			 SERIAL PRIMARY KEY,
	time_stamp 	 TIMESTAMPTZ,
	name		 VARCHAR(255),
	distance 	 INTEGER CHECK (distance >= 0),
	gallons_used INTEGER CHECK (gallons_used >= 0),
	cost_of_fuel INTEGER CHECK (cost_of_fuel >= 0),
	flight 	 VARCHAR(255),

	unique(name, time_stamp)
);

-- For storage
CREATE TABLE twitter_sources (
	twitter_id 	  VARCHAR(255) PRIMARY KEY,
	username 	  VARCHAR(50),
	last_tweet_id VARCHAR(255)
);

CREATE TABLE new_trip_log (
	id 		   SERIAL PRIMARY KEY,
	time_stamp TIMESTAMPTZ,
	log 	   JSON NOT NULL
);