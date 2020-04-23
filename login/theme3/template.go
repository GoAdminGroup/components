package theme3

var List = map[string]string{"login": `{{define "login"}}
    <!DOCTYPE html>
    <!--[if lt IE 7]>
    <html class="no-js lt-ie9 lt-ie8 lt-ie7">
    <![endif]-->
    <!--[if IE 7]>
    <html class="no-js lt-ie9 lt-ie8">
    <![endif]-->
    <!--[if IE 8]>
    <html class="no-js lt-ie9">
    <![endif]-->
    <!--[if gt IE 8]><!-->
    <html class="no-js">
    <!--<![endif]-->
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>{{.Title}}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.css"}}">

        <!--[if lt IE 9]>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/respond.min.js"}}"></script>
        <![endif]-->

    </head>
    <body data-gr-c-s-loaded="true" style="">
    <div class="page login-page">
        <div class="container d-flex align-items-center">
            <div class="form-holder has-shadow">
                <div class="row">

                    <div class="col-lg-6">
                        <div class="info d-flex align-items-center">
                            <div class="content">
                                <div class="logo">
                                    <h1>{{lang "Welcome Login"}}</h1>
                                </div>
                                <p>{{.Title}}</p>
                            </div>
                        </div>
                    </div>

                    <div class="col-lg-6 bg-white">
                        <div class="form d-flex align-items-center">
                            <div class="content">
                                <form method="post" action="##" onsubmit="return false" class="form-validate" novalidate="novalidate">
                                    <div class="form-group">
                                        <input id="username" type="text" name="userName" required=""
                                               placeholder="{{lang "username"}}" class="input-material is-invalid" aria-invalid="true"
                                               aria-describedby="login-username-error">
                                        <div id="username-error" class="is-invalid invalid-feedback"
                                             style="display: none;">{{lang "input"}} {{lang "username"}}
                                        </div>
                                    </div>
                                    <div class="form-group">
                                        <input id="password" type="password" name="passWord" required=""
                                               placeholder="{{lang "password"}}" class="input-material is-invalid"
                                               aria-invalid="true" aria-describedby="login-password-error">
                                        <div id="password-error" class="is-invalid invalid-feedback"
                                             style="display: none;">{{lang "input"}} {{lang "password"}}
                                        </div>
                                    </div>
                                    {% if .CaptchaDigits %}
                                    <div class="form-group">
                                        <div class="row" style="margin: 0px;">
                                            <div class="col-lg-7">
                                                <input type="text" required="" class="input-material form-control" placeholder="{{lang "captcha"}}" id="captcha">
                                                <div id="captcha-error" class="is-invalid invalid-feedback"
                                                     style="display: none;">{{lang "input"}} {{lang "captcha"}}
                                                </div>
                                            </div>
                                            <div class="col-lg-5">
                                                <img class="captcha" src="{% .CaptchaImgSrc %}" alt="" width="120" height="38">
                                            </div>
                                        </div>
                                        <input type="hidden" value="{% .CaptchaID %}" id="captcha_id">
                                    </div>
                                    {% end %}
                                    <button id="login" type="submit" class="btn btn-primary" onclick="submitData()">{{lang "login"}}</button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    </body>

    <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.js"}}"></script>


    {% if .TencentWaterProofWallData.AppID  %}
    <script src="https://ssl.captcha.qq.com/TCaptcha.js"></script>
    {% end %}

    <script>

        {% if .TencentWaterProofWallData.AppID  %}

        let captcha = new TencentCaptcha("{% .TencentWaterProofWallData.AppID %}", function (res) {
            // res（用户主动关闭验证码）= {ret: 2, ticket: null}
            // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
            if (res.ret === 0) {
                $.ajax({
                    dataType: 'json',
                    type: 'POST',
                    url: '{{.UrlPrefix}}/signin',
                    async: 'true',
                    data: {
                        'username': $("#username").val(),
                        'password': $("#password").val(),
                        'token': res.ticket+","+res.randstr
                    },
                    success: function (data) {
                        location.href = data.data.url
                    },
                    error: function (data) {
                        alert(data.responseJSON.msg);
                    }
                });
            } else {
                alert(data.data.msg);
            }
        }, {});

        {% end %}

        {% if .CaptchaDigits %}

        $(".captcha").on("click",function(){
            location.reload();
        });

        {% end %}

        function submitData() {
            {% if .TencentWaterProofWallData.AppID  %}
            captcha.show();
            {% else  %}
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '{{.UrlPrefix}}/signin',
                async: 'true',
                data: {
                    'username': $("#username").val(),
                    {% if .CaptchaDigits %}
                        'token': $("#captcha").val() + "," + $("#captcha_id").val(),
                    {% end %}
                    'password': $("#password").val()
            },
            success: function (data) {
                location.href = data.data.url
            },
            error: function (data) {
                alert(data.responseJSON.msg);
				location.reload();
            }
        });
            {% end %}
        }
    </script>

    </body>
    </html>
{{end}}`}
