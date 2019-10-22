-- Create tables.

DROP TABLE IF EXISTS "balancers" CASCADE;
CREATE TABLE "balancers"
(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50)
);

DROP TABLE IF EXISTS "machines" CASCADE;
CREATE TABLE "machines"
(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50),
    "is_working" INT
);

DROP TABLE IF EXISTS "relations";
CREATE TABLE "relations"
(
	"id" SERIAL PRIMARY KEY,
	"id_balancer" INT NOT NULL,
	"id_machine" INT NOT NULL,
	FOREIGN KEY ("id_balancer") REFERENCES "balancers"(id),
    FOREIGN KEY ("id_machine") REFERENCES "machines"(id)
);

-- Insert demo data.
INSERT INTO "balancers" (name) VALUES ('balancer1');
INSERT INTO "balancers" (name) VALUES ('balancer2');

INSERT INTO "machines" (name, is_working) VALUES ('machine1', 1);
INSERT INTO "machines" (name, is_working) VALUES ('machine2', 0);
INSERT INTO "machines" (name, is_working) VALUES ('machine3', 1);
INSERT INTO "machines" (name, is_working) VALUES ('machine4', 0);
INSERT INTO "machines" (name, is_working) VALUES ('machine5', 1);

INSERT INTO "relations" (id_balancer, id_machine) VALUES (2, 1);
INSERT INTO "relations" (id_balancer, id_machine) VALUES (1, 2);
INSERT INTO "relations" (id_balancer, id_machine) VALUES (2, 3);
INSERT INTO "relations" (id_balancer, id_machine) VALUES (1, 4);
INSERT INTO "relations" (id_balancer, id_machine) VALUES (2, 5);
