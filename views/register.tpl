{{ template "base.tpl" . }}

{{ define "Title" }}
    Pratl - Register
{{ end }}

{{ define "Content" }}
    <br>
    <div class="container">
        <div class="row">
            <div class="col-md-4 col-md-offset-4">
                <h1>Register</h1>
                <p class="text-muted">
                    Already have an account?
                    <a href="{{ urlfor "UserController.Login" }}">Login here.</a>
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
                        <input class="form-control" name="Name" value="{{ .Form.Name }}" placeholder="Name">
                    </p>
                    <p>
                        <input type="password" class="form-control" name="Password" value="{{ .Form.Password }}" placeholder="Password">
                    </p>
                    <p>
                        <input type="password" class="form-control" name="Password2" value="{{ .Form.Password2 }}" placeholder="Confirm Password">
                    </p>
                    <button type="submit" class="btn btn-default">
                        Register
                    </button>
                </form>
            </div>
        </div>
    </div>
{{ end }}
