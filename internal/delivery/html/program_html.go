package html
 const ProgramHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Programs Table</title>
	<style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      padding: 20px;
      color: #333;
    }

    h1,
    h2 {
      color: #444;
    }

    form {
      background-color: #fff;
      padding: 20px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
      margin-bottom: 20px;
    }

    label {
      display: block;
      margin-top: 10px;
      color: #666;
    }

    input,
    select,
    button {
      width: 100%;
      padding: 10px;
      margin-top: 5px;
      margin-bottom: 20px;
      border-radius: 5px;
      border: 1px solid #ddd;
      box-sizing: border-box; /* Makes sure the padding doesn't affect the total width */
    }

    button {
      background-color: #5cb85c;
      color: white;
      border: none;
      cursor: pointer;
    }

    button:hover {
      background-color: #4cae4c;
    }

    table {
      width: 100%;
      border-collapse: collapse;
      background-color: #fff;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    th,
    td {
      padding: 10px;
      border: 1px solid #ddd;
      text-align: left;
    }

    th {
      background-color: #f8f8f8;
    }

    tr:nth-child(even) {
      background-color: #f2f2f2;
    }

    /* Responsive Table */
    @media screen and (max-width: 600px) {
      table {
        border: 0;
      }

      table caption {
        font-size: 1.3em;
      }

      table thead {
        display: none;
      }

      table tr {
        margin-bottom: 10px;
        display: block;
        border-bottom: 2px solid #ddd;
      }

      table td {
        display: block;
        text-align: right;
        font-size: 0.8em;
        border-bottom: 1px dotted #ccc;
      }

      table td:last-child {
        border-bottom: 0;
      }

      table td::before {
        content: attr(data-label);
        float: left;
        font-weight: bold;
        text-transform: uppercase;
      }

      table td:last-child {
        border-bottom: 0;
      }
    }
  </style>
</head>
<body>
    <h1>Create Program</h1>
    <form method="POST" action="https://curifyapp.up.railway.app/save/program/upload" enctype="multipart/form-data">
        <label for="ageUp">Age Up:</label>
        <input type="number" name="AgeUp" required><br>
        <br>
        <label for="ageDown">Age Down:</label>
        <input type="number" name="AgeDown" required><br>
        <br>
        <label for="bmiUp">BMI Up:</label>
        <input type="number" step="0.01" name="BMIUp" required><br>
        <br>
        <label for="bmiDown">BMI Down:</label>
        <input type="number" step="0.01" name="BMIDown" required><br>
        <br>
        <label for="type">Program Type:</label>
        <select name="Type" required>
            <option value="weight_loss">weight_loss</option>
            <option value="stress_work">stress_work</option>
            <!-- Add more options as needed -->
        </select><br>

        <button type="submit">Submit</button>
    </form>

    <h2>Programs Table</h2>
    <table border="1" id="programsTable">
        <tr>
            <th>ID</th>
            <th>Age Up</th>
            <th>Age Down</th>
            <th>BMI Up</th>
            <th>BMI Down</th>
            <th>Program Type</th>
            <th>Pro Type</th>
        </tr>
    </table>

    <script>
        var jsonData = {{.ProgramsJSON}};

        function populateTable(data) {
            var table = document.getElementById("programsTable");

            data.forEach(function(program) {
                var row = table.insertRow();
                row.insertCell(0).textContent = program.id;
                row.insertCell(1).textContent = program.ageUp;
                row.insertCell(2).textContent = program.ageDown;
                row.insertCell(3).textContent = program.bmiUp;
                row.insertCell(4).textContent = program.bmiDown;
                row.insertCell(5).textContent = program.type;
                row.insertCell(6).textContent = program.proType;
            });
        }

        // Call the function with the JSON data
        populateTable(jsonData);
    </script>
</body>
</html>

`