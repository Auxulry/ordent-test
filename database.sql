CREATE TABLE "users" (
  "created_at" timestamp(0) NOT NULL,
  "updated_at" timestamp(0) NOT NULL,
  "deleted_at" timestamp(0) NULL,
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(100) NOT NULL,
  "balance" numeric NOT NULL DEFAULT 0.00
);

CREATE TABLE "products" (
  "created_at" timestamp(0) NOT NULL,
  "updated_at" timestamp(0) NOT NULL,
  "deleted_at" timestamp(0) NULL,
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "sku" character varying(5) NOT NULL,
  "name" character varying NOT NULL,
  "description" text NOT NULL,
  "price" numeric NOT NULL
);

INSERT INTO "products" ("sku", "name", "description", "price") VALUES
('TB001', 'Black T-Shirt', 'Black T-shirt all size', 120000),
('BB001', 'Black Bag', 'Black Bag all size', 80000),
('HR001', 'Red Hoodie', 'Red Hoodie all size', 150000);

CREATE TABLE cart_sessions (
    "created_at" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp(0) NULL,
    "id" SERIAL NOT NULL,
    PRIMARY KEY("id"),
    "user_id" integer NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
    total numeric NOT NULL
);

CREATE TABLE carts (
    "created_at" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp(0) NULL,
    "id" SERIAL NOT NULL,
    PRIMARY KEY("id"),
    "cart_session_id" integer NOT NULL,
    CONSTRAINT fk_cart_session FOREIGN KEY ("cart_session_id") REFERENCES "cart_sessions" ("id"),
    "product_id" integer NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products" ("id"),
    quantity int NOT NULL
);