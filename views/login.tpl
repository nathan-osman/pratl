{{ template "base.tpl" . }}

{{ define "Title" }}
    Pratl - Login
{{ end }}

{{ define "Content" }}
    <br>
    <div class="container">
        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <h1>Login</h1>
                <p class="text-muted">
                    Don't have an account?
                    <a href="{{ urlfor "UserController.Register" }}">Register here.</a>
                </p>
                {{if .Error }}
                    <div class="alert alert-danger">
                        {{ .Error }}
                    </div>
                {{ end }}
                <form method="post">
                    <br>
                    <div class="form-group">
                        <label for="email">Email</label>
                        <input type="email" class="form-control" id="email" name="Email" value="{{ .Form.Email }}">
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input type="password" class="form-control" id="password" name="Password" value="{{ .Form.Password }}">
                    </div>
                    <br>
                    <button type="submit" class="btn btn-default btn-lg">
                        Login
                    </button>
                </form>
            </div>
        </div>
    </div>
{{ end }}
