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
        <form action="/db" method="post">
            <label style="color: white;">{{.}}</label><br><br>
            <label style="color: white;">Stations:</label><input type="text" name="Stations">&emsp;<br><br>
            <input type="submit" value="Search in DB">
        </form>
    </body>
</html>