<html>
    <head>
    <title></title>
    <style>
        body{
            background-color: #000000e7; 
        }
    </style>
    </head>
    <body>
        <form action="/wl" method="post">
            <label style="color: white;">{{.}}</label><br><br>
            <label style="color: white;">Stations:</label><input type="text" name="Stations">&emsp;
            <input type="checkbox" name="DBFlag" value="AddToDB" style="color: white;"> <label style="color: white;">Add to DB</label><br><br>
            <input type="submit" value="Get Workload">
        </form>
    </body>
</html>