<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <title>{{ template "Title" . }}</title>

        <!-- Favicons -->
        <link rel="icon" href="/static/img/favicon.png">
        <link rel="apple-touch-icon" href="/static/img/apple-touch-icon.png">

        <!-- Stylesheets -->
        <link rel="stylesheet" href="/static/css/bootstrap.min.css">
        <link rel="stylesheet" href="/static/css/style.css">
        {{ template "Stylesheets" . }}
    </head>
    <body>
        <!-- Content -->
        {{ template "Content" . }}

        <!-- Scripts -->
        <script src="/static/js/jquery.min.js"></script>
        <script src="/static/js/bootstrap.min.js"></script>
        {{ template "Scripts" . }}
    </body>
</html>

{{ define "Stylesheets" }}{{ end }}
{{ define "Scripts" }}{{ end }}
