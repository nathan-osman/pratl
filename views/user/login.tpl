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
                    <p>
                        <input type="email" class="form-control" name="Email" value="{{ .Form.Email }}" placeholder="Email">
                    </p>
                    <p>
                        <input type="password" class="form-control" name="Password" value="{{ .Form.Password }}" placeholder="Password">
                    </p>
                    <button type="submit" class="btn btn-default">
                        Login
                    </button>
                </form>
            </div>
        </div>
    </div>
{{ end }}
