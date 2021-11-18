<html>
    <head>
    <title>Ascii-Art</title>
    </head>
    <body>
    <h1>Ascii-Art Web</h1>
        <form action="/ascii-art" method="post">
            <p>Please select your template:</p>
            <input type="radio" id="banner1" name="banner" value="shadow.txt">
            <label for="banner1">Shadow</label><br>
            <input type="radio" id="banner2" name="banner" value="standard.txt">
            <label for="banner2">Standard</label><br>  
            <input type="radio" id="banner3" name="banner" value="thinkertoy.txt">
            <label for="banner3">Thinkertoy</label><br><br>
            Enter word:<br><br>
            <textarea  id="word" name="word"></textarea><br><br>
            <input type="submit" value="Submit">
        </form>
    </body>
</html>