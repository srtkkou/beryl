doctype 5
html.has-navbar-fixed-top[lang="ja"]
  head
    meta[charset="UTF-8"]
    meta[http-equiv="X-UA-Compatible"][content="IE=edge"]
    meta[name="viewport"][content="width=device-width,initial-scale=1.0"]
    meta[name="robots"][content="noindex, nofollow"]
    title
      #{$.I18n(".layout.Title")} 
      {{template title}}
    // アイコン
    link[rel="icon"][href="/image/favicon.ico"]
    link[rel="apple-touch-icon"][sizes="180x180"]
      [href="/image/apple-touch-icon.png"]
    // CSSライブラリ
    link[rel="stylesheet"]
      [href="https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css"]
    {{template css}}
  body
    section.section
      div.container
        // メッセージ
        {{template messages}}
        // メインコンテンツ
        {{template content}}
    footer.footer
      div.content
        ul
          li
            | Copyright &copy; 2021 Coexe Co. ,Ltd. All Rights Reserved.
          li
            | Ver.#{Version}
    // JavaScriptライブラリ
    script[src="https://kit.fontawesome.com/0784eb9fea.js"]
      [crossorigin="anonymous"]
    script[src="https://code.jquery.com/jquery-3.5.0.min.js"]
      [integrity="sha256-xNzN2a4ltkB44Mc/Jz3pT4iU1cmeR0FkXs4pru/JxaQ="]
      [crossorigin="anonymous"]
    script[src="/js/app.js"]
    {{template js}}
