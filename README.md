# Go-Paper-API

The building project was created by the famous architects  Kirill Illarionov and Alexandr Dolgavin, who combined two styles in the structure: classic and Ukrainian Baroque. 


# How to run:
## Step 1 (install postgresql): 
- sudo apt-get update
- sudo apt-get install postgresql postgresql-contrib
## Step 2 (create user and db): 
- sudo -i -u postgres 
- sudo -u postgres createdb docker
- psql
- CREATE USER docker WITH PASSWORD 'docker';
## Step 3 (init database):
- copy all from init.sql from internal/initDatabase and paste it in database
## Step 4 (run):
- go run main.go (You can also use nohup for run server like daemon on a VPS)
