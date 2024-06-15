<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>Site CTF</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f0f0f0;
            color: #333;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            text-align: center;
        }
        h1 {
            color: #4CAF50;
        }
        p {
            margin: 20px 0;
        }
        form {
            margin-top: 20px;
        }
        select {
            padding: 10px;
            margin-right: 10px;
            border-radius: 4px;
            border: 1px solid #ccc;
        }
        button {
            padding: 10px 15px;
            background-color: #4CAF50;
            color: #fff;
            border: none;
            border-radius: 4px;
            cursor: pointer;
        }
        button:hover {
            background-color: #45a049;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Bienvenue sur le Site CTF</h1>
        <p>Changer de langue :</p>
        <form method="GET" action="lang.php">
            <select name="lang">
                <option value="en.php">English</option>
                <option value="fr.php">Français</option>
            </select>
            <button type="submit">Changer de langue</button>
        </form>
    </div>
</body>
</html>