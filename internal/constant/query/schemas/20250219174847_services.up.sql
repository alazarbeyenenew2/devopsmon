CREATE TABLE services ( 
         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
          name VARCHAR NOT NULL,
          description verchar not null
);