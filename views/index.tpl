{{ template "base.tpl" . }}

{{ define "Title" }}
    Pratl - Chat Made Simple
{{ end }}

{{ define "Content" }}
    <div class="container">
        <br><br>
        <div class="row">
            <div class="col-md-4 col-md-offset-2">
                <img src="/static/img/window.svg">
            </div>
            <div class="col-md-4">
                <h1>
                    Pratl
                </h1>
                <p class="text-muted">
                    Chat Made Simple
                </p>
                <hr>
                <ul>
                    <li>Conversations organized into channels</li>
                    <li>Distraction-free interface</li>
                    <li>Responsive design</li>
                    <li>Incredibly easy to deploy</li>
                    <li>Scales to hundreds of users</li>
                    <li>Completely free and open-source</li>
                </ul>
                <hr>
                <a href="/login" class="btn btn-default btn-lg">
                    Login
                </a>
                <a href="/register" class="btn btn-default btn-lg">
                    Register
                </a>
            </div>
        </div>
        <br><br>
        <p class="text-center small text-muted">
            Copyright 2016 &mdash; Nathan Osman
        </p>
    </div>
{{ end }}
