package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	EOF     = "EOF"

	Indent       = "  "
	LF           = "\n"
	Comment      = "#"
	PrintComment = "#+"

	Period    = "."
	LBracket  = "["
	RBracket  = "]"
	Quote     = `"`
	Assign    = "="
	Code      = "-"
	PrintCode = "+"

	Html  = "Html"  // HTML自体のタグ
	Head  = "Head"  // ページ自体のメタ情報をまとめるタグ
	Title = "Title" // ページ自体の題名のタグ
	Link  = "Link"  // 外部ファイルとのリンクのタグ
	/*
		meta 	ページ自体のメタ情報のタグ
		style 	主にCSSをページ自体に書き込むタグ
		script 	主にJavaScriptをページに埋め込むタグ
		noscript 	主にJavaScriptの代わりに表示するコンテンツのタグ
		body 	ブラウザに表示する内容全体のタグ
		h1 	見出しのタグ1
		h2 	見出しのタグ2
		h3 	見出しのタグ3
		h4 	見出しのタグ4
		h5 	見出しのタグ5
		h6 	見出しのタグ6
		p 	段落のタグ
		hr 	段落と段落の区切りのタグ
		blockquote 	引用するコンテンツのタグ
		ol 	順番の決まったリストのタグ
		ul 	順番の決まってないリストのタグ
		li 	リストの内容のタグ
		dl 	項目に関する記述リストのタグ
		dt 	項目に関する記述リストの項目のタグ
		dd 	項目に関する記述リストの記述のタグ
		div 	特定の意味を持たないブロックのタグ
		a 	リンクのタグ
		strong 	文章内の重要性の高いテキストのタグ
		span 	文章内の特定の意味を持たないテキストのタグ
		br 	文章内の改行のタグ
		img 	写真や画像のタグ
		table 	テーブルのタグ
		tr 	テーブル内の行のタグ
		td 	テーブル内の行の内容のタグ
		th 	テーブル内の行の見出しのタグ
		form 	フォームのタグ
		label 	フォーム内の項目にラベルを付けるタグ
		input 	フォーム内の入力項目のタグ
		select 	フォーム内のメニューのタグ
		option 	フォーム内のメニューの選択肢のタグ
		textarea 	フォーム内の文章の入力項目のタグ
		base 	ページ内の相対URLの基準になるURLのタグ
		section 	見出しと内容をまとめるセクションのタグ
		nav 	リンクをまとめるセクションのタグ
		article 	単独で扱える記事のようなセクションのタグ
		aside 	本筋から逸れる解説などのセクションのタグ
		hgroup 	二つ以上の見出しをまとめるタグ
		header 	コンテンツのヘッダー情報をまとめるタグ
		footer 	コンテンツのフッター情報をまとめるタグ
		address 	コンテンツの作者の連絡先のタグ
		pre 	書いたままのテキストを表示するタグ
		figure 	図表や画像とその説明をまとめるタグ
		figcaption 	図表や画像につける説明のタグ
		em 	文章内の特に強調するテキストのタグ
		small 	文章内の補足や細目のテキストのタグ
		cite 	文章内の作品のタイトルのタグ
		q 	文章内の引用部分のテキストのタグ
		abbr 	文章内の略称や頭字語のテキストのタグ
		time 	文章内の日時のタグ
		code 	文章内のプログラムなどのコードのタグ
		ins 	後から追記するコンテンツのタグ
		del 	後から削除するコンテンツのタグ
		iframe 	ページ内にフレームを埋め込むタグ
		embed 	ページ内に主にFlash等を埋め込むタグ
		object 	ページ内に画像や音声、動画を埋め込むタグ
		param 	objectタグのコンテンツのパラメータのタグ
		video 	ページ内に動画ファイルを埋め込むタグ
		audio 	ページ内に音声ファイルを埋め込むタグ
		canvas 	ページ内にcanvasを埋め込むタグ
		caption 	テーブルに付けるキャプションのタグ
		tbody 	テーブル内のボディ部分のデータをまとめるタグ
		thead 	テーブル内のヘッダー部分の行のタグ
		tfoot 	テーブル内のフッター部分の行のタグ
		fieldset 	フォーム内の内容をまとめるタグ
		legend 	fieldsetタグにつけるタイトルのタグ
		s 	文章内の打ち消し線のついたテキストのタグ
		dfn 	文章内の定義する用語のテキストのタグ
		var 	文章内の変数のテキストのタグ
		samp 	文章内のプログラムで出力するテキストのタグ
		kbd 	文章内の入力テキストかキーボードのキーのタグ
		sub 	文章内の上下に付くテキストのタグ
		sup 	文章内の下上に付くテキストのタグ
		i 	文章内の慣習上斜体にするテキストのタグ
		b 	文章内の慣習上太字にするテキストのタグ
		u 	文章内の慣習上下線を引くテキストのタグ
		mark 	文章内の目立たせたテキストのタグ
		ruby 	文章内のルビを付けるテキストのタグ
		rt 	ルビとして付くテキストのタグ
		rp 	ルビが表示されない場合のカッコのタグ
		bdi 	文章内の書く方向を独立させるテキストのタグ
		bdo 	文章内の書く方向を決めるテキストのタグ
		wbr 	文章内の改行可能な位置のタグ
		source 	ページ内に埋め込む動画や音声を指定するタグ
		track 	動画や音声と同期するテキストトラックのタグ
		map 	クリックできるイメージマップのタグ
		area 	mapのクリックする場所を決めるタグ
		colgroup 	colタグをまとめるタグ
		col 	テーブル内の縦列をまとめるタグ
		button 	フォーム内のボタンのタグ
		datalist 	フォーム内の入力項目に選択肢をつけるタグ
		optgroup 	フォーム内の複数のoptionタグをまとめるタグ
		keygen 	フォーム内で暗号になるキーを作るタグ
		output 	フォーム内の計算結果の出力のタグ
		progress 	処理の進捗状況などのゲージのタグ
		meter 	予め数値の決まったゲージのタグ
		detail 	クリックすると見られる詳細情報のタグ
		summary 	detailタグにつけるラベルのタグ
		command 	右クリックメニューやキーボード操作の内容のタグ
		menu 	ツールバーや右クリックメニューを埋め込むタグ
	*/
)
