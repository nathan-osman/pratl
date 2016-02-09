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
                    <br>
                    <div class="form-group">
                        <label for="email">Email</label>
                        <input type="email" class="form-control" id="email" name="Email" value="{{ .Form.Email }}">
                    </div>
                    <div class="form-group">
                        <label for="name">Name</label>
                        <input class="form-control" id="name" name="Name" value="{{ .Form.Name }}">
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input type="password" class="form-control" id="password" name="Password" value="{{ .Form.Password }}">
                    </div>
                    <div class="form-group">
                        <label for="password2">Confirm Password</label>
                        <input type="password" class="form-control" id="password2" name="Password2" value="{{ .Form.Password2 }}">
                    </div>
                    <br>
                    <button type="submit" class="btn btn-primary">
                        Register
                    </button>
                </form>
            </div>
        </div>
    </div>
{{ end }}
