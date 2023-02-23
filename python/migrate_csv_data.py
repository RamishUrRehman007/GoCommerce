import psycopg2
import csv

connection = psycopg2.connect("host='postgres' port='5432' dbname='go_commerce' user='user' password='password'")
mycursor = connection.cursor()

def insert_customer_companies():
    with open('Test task - Postgres - customer_companies.csv', newline='',  encoding="utf8") as csvfile:
        spamreader = csv.reader(csvfile)
        for row in spamreader:
            # Prepare SQL query to INSERT a record into the database.
            sql = "INSERT INTO customer_companies(id, company_name) VALUES ('%s', '%s');" % (row[0], row[1])
            print(sql)
            try:
               # Execute the SQL command
               mycursor.execute(sql)
               # Commit your changes in the database
               connection.commit()
            except:
               # Rollback in case there is any error
               connection.rollback()


def insert_customers():
    with open('Test task - Postgres - customers.csv', newline='',  encoding="utf8") as csvfile:
        spamreader = csv.reader(csvfile)
        for row in spamreader:
            # Prepare SQL query to INSERT a record into the database.
            sql = "INSERT INTO customers (user_id, login, password, name, company_id, credit_cards) VALUES ('%s','%s','%s','%s','%s','%s');" % (row[0], row[1], row[2], row[3], row[4], row[5])
            print(sql)
            try:
               # Execute the SQL command
               mycursor.execute(sql)
               # Commit your changes in the database
               connection.commit()
            except:
               # Rollback in case there is any error
               connection.rollback()


def insert_orders():
    with open('Test task - Postgres - orders.csv', newline='',  encoding="utf8") as csvfile:
        next(csvfile)
        spamreader = csv.reader(csvfile)
        for row in spamreader:
            # Prepare SQL query to INSERT a record into the database.
            sql = "INSERT INTO orders(id, created_at, order_name, customer_id) VALUES ('%s','%s', '%s', '%s' );" % (row[0], row[1], row[2], row[3])
            print(sql)
            try:
               # Execute the SQL command
               mycursor.execute(sql)
               # Commit your changes in the database
               connection.commit()
            except:
               # Rollback in case there is any error
               connection.rollback()


def inser_order_items():
    with open('Test task - Postgres - order_items.csv', newline='',  encoding="utf8") as csvfile:
        next(csvfile)
        spamreader = csv.reader(csvfile)
        for row in spamreader:
            # Prepare SQL query to INSERT a record into the database.
            sql = "INSERT INTO order_items(id, order_id, price_per_unit, quantity, product) VALUES ('%s','%s', '%s', '%s', '%s');" % (row[0], row[1], row[2], row[3], row[4])
            print(sql)
            try:
               # Execute the SQL command
               mycursor.execute(sql)
               # Commit your changes in the database
               connection.commit()
            except:
               # Rollback in case there is any error
               connection.rollback()


def inser_deliveries():
    with open('Test task - Postgres - deliveries.csv', newline='',  encoding="utf8") as csvfile:
        next(csvfile)
        spamreader = csv.reader(csvfile)
        for row in spamreader:
            # Prepare SQL query to INSERT a record into the database.
            sql = "INSERT INTO deliveries(id, order_item_id, delivered_quantity) VALUES ('%s','%s', '%s');" % (row[0], row[1], row[2])
            print(sql)
            try:
               # Execute the SQL command
               mycursor.execute(sql)
               # Commit your changes in the database
               connection.commit()
            except:
               # Rollback in case there is any error
               connection.rollback()


insert_customer_companies()
insert_customers()
insert_orders()
inser_order_items()
inser_deliveries()
connection.close()