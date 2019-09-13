package static

//BustData contains the format of the json file
type BustData struct {
	Url    string  `json:"url,omitempty"`
	Method string  `json:"method,omitempty"`
	Status string  `json:"status,omitempty"`
	Error  *string `json:"error,omitempty"`
	Extra  *string `json:"extra,omitempty"`
}

var HtmlCode = `
<!DOCTYPE html>
<html>
<head>
<title>Web Buster</title>
</head>
<body>
    <link rel="stylesheet" href="https://cdn.datatables.net/1.10.19/css/dataTables.bootstrap4.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.1/css/bootstrap.css">

    <script src="https://code.jquery.com/jquery-3.3.1.js"></script>
    <script src="https://cdn.datatables.net/1.10.19/js/jquery.dataTables.min.js"></script>
    <script src="https://cdn.datatables.net/1.10.19/js/dataTables.bootstrap4.min.js"></script>


<br>
<h1 align='center'>WebBuster</h1>
<br>

<div class="container">
        <table id="myTable"class="table table-striped table-bordered" style="width:100%">
            <thead class="thead-dark">
                <tr>
                    <th scope="col">URL</th>
                    <th scope="col">Method</th>
                    <th scope="col">Status</th>
                    <th scope="col">Error</th>
                    <th scope="col">Extra</th>
                </tr>
            </thead>
        </table>
    </div>

<script>
    $(document).ready(function() {
        $('#myTable').DataTable( {
        "responsive": true,
        "aaData":{{.}},
        "aoColumns":
        [
			{"mDataProp": "url", "defaultContent": "n/a", className: "text-center"},
            {"mDataProp": "method", "defaultContent": "n/a", className: "text-center"},
            {"mDataProp": "status", "defaultContent": "n/a", className: "text-center"},
            {"mDataProp": "error", "defaultContent": "n/a", className: "text-center"},
            {"mDataProp": "extra", "defaultContent": "n/a", className: "text-center"}
        ]
        } );
    } );

    </script>
</body>
</html>
`
