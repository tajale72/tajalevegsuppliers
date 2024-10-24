# my one app is an app which creates bills, and save the bill information in both postrges and mongoDB.
# It also has product list endpoint which returns all the records in the database.


TRUNCATE TABLE bill_details RESTART IDENTITY;

nohup ./build_deploy.sh > output.log 2>&1 &

sudo lsof -i :8080
