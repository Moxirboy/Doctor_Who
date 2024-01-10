package html

const ExerciseHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>exercise  Table</title>

	<style>
      body {
        background: #e9ecef;
        font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
        padding: 20px;
        line-height: 1.5;
      }

      h1,
      h2 {
        color: #333;
      }

      /* Form styling */
      form {
        background: #fff;
        padding: 20px;
        border-radius: 5px;
        margin-bottom: 20px;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
      }

      label {
        display: block;
        margin-top: 10px;
        color: #333;
      }

      input[type="number"],
      input[type="text"] {
        width: 100%;
        padding: 8px;
        margin-top: 5px;
        margin-bottom: 20px;
        border: 1px solid #ced4da;
        border-radius: 4px;
      }

      button {
        background-color: #5cb85c;
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
      }

      button:hover {
        background-color: #5cb85c;
      }

      /* Table styling */
      table {
        width: 100%;
        border-collapse: collapse;
        background-color: #fff;
        margin-top: 20px;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
      }

      th,
      td {
        padding: 10px;
        border: 1px solid #dee2e6;
        text-align: left;
      }

      th {
        background-color: #f8f9fa;
      }

      tr:nth-child(even) {
        background-color: #f2f2f2;
      }

      /* Removes default browser styling for <br> */
      br {
        display: none;
      }
    </style>
</head>
<body>
    <h1>Create exercise </h1>
    <form method="POST" action="https://curifyapp.up.railway.app/save/exercise/upload">
        <label for="program_id">Program ID:</label>
        <input type="number" name="program_id" required><br>
        <br>
        <label for="name">Name:</label>
        <input type="text" name="name" required><br>
        <br>
        <label for="info">Info:</label>
        <input type="text" name="info" required><br>
        <br>
        <label for="link_to_video">Link to Video:</label>
        <input type="text" name="link_to_video" required><br>
        <br>
        <button type="submit">Submit</button>
    </form>

    <h2>Exercise Table</h2>
        <table border="1" id="medicalProgramsTable">
            <tr>
                <th>ID</th>
                <th>Program ID</th>
                <th>Name</th>
                <th>Info</th>
                <th>Link to Video</th>
            </tr>
        </table>

        <script>
            var medicalProgramsData = {{.MedicalProgramsJSON}};

            function populateMedicalProgramsTable(data) {
                var table = document.getElementById("medicalProgramsTable");

                data.forEach(function(program) {
                    var row = table.insertRow();
                    row.insertCell(0).textContent = program.id;
                    row.insertCell(1).textContent = program.programId;
                    row.insertCell(2).textContent = program.name;
                    row.insertCell(3).textContent = program.info;
                    row.insertCell(4).textContent = program.link;
                });
            }

            // Call the function with the JSON data
            populateMedicalProgramsTable(medicalProgramsData);
        </script>
</body>
</html>

`

