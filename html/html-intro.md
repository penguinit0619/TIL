# HTML
> Intro

<p>사실 회사에서는 Java만 주구장창 하다보니깐 마크업 관련 언어는 전혀 건들일 일이 없다.  이번에 Hugo를 통해서 블로그를 만들려고 마크업작업을 하였는데 이게 왠일인지 너무 어려웠다. 내심 HTML이니깐 대충해도 하겠지라는 생각이었는데... 디자인이라던가 완전 문외한이라 저번주에 책을 샀다.</p>

> [HTML CSS Design Recipe](http://www.yes24.com/24/goods/66282054)

<p>일단 HTML이나 CSS쪽 TIL은 이 책에 있는 예제나 내용들로 적을 예정이다. HTML에 대해서 깊게 파기 보다는 테크닉위주로 기록하려고 한다. (책내용도 테크닉위주)</p>

그래도 잡설만 하면 좀 그러니 HTML에서 필수적으로 명시해주어야 하는 태그들에 대해서 Intro에 적으려고한다.

### DOCTYPE
```html
<!DOCTYPE html> <!-- DOCTYPE 선언 -->
```
HTML 첫번 째 줄에는 반드시 DOCTYPE 선언을 써야한다. 이 DOCTYPE선언은 현재 내가 작성하려는 HTML페이지가 어떤 버전을 기준으로 작성 되어 있는지를 표시하는 것인데 위에 작성한 DOCTYPE은 최신의 DOCTYPE을 가리키는 것으로 문제가 되지 않는다. HTML5.x 이전에 4.x나 XHTML 시절에는 저것보다 복잡한 DOCTYPE을 명시해야 했던 것으로 기억하고 있다.

### HTML lang속성
```html
<html lang="ko"> <!-- html 언어속성 적용 -->
```
이 부분은 여러번 보았는데 그 의미를 자세하게 모르고 있었던 부분중 하나다. 이 lang 속성을 명시하지 않는다고 해서 페이지의 콘텐츠에 뭔가 영향이 가지는 않지만 검색엔진에서 번역기능이 html 태그의 lang속성을 보고 번역유무를 판단한다. 예를 들면 한글 브라우저에서 영어 페이지를 보면 `이 사이트를 번역하시겠습니까?` 라는 메시지창은 html lang 속성이 en으로 설정되어 있기에 나오는 메세지이다.

아래는 주로사용되는 언어코드이다.


| 언어코드 | 언어   |
| ---- | ---- |
| ko   | 한국어  |
| en   | 영어   |
| zh   | 중국어  |
| ja   | 일본어  |
| es   | 스페인어 |
| de   | 독일어  |
| fr   | 프랑스어 |


### Character Set
```html
<meta charset="UTF-8"> <!-- Character Set를 UTF-8로 설정 -->
```
위와 같이 선언하면 현재 작성하고 있는 문서는 UTF-8의 Character Set을 따른다는 의미이다. 내 경우에는 UTF-8을 무의식적으로 쓰긴하는데 Character Set에 대해서 조금더 알아보자면

- 컴퓨터에 표시되는 모든 문자에는 문자코드라고 불리는 ID번호가 있는데 Character Set마다 각 주어진 ID번호가 다르게 저장된다. 예를 들어 "가" 라는 글자는 UTF-8에서 ID번호는 `EAB080`이지만 EUC-KR에서는 `BOA1`이다. 이렇게 되면 Character Set이 맞지 않았을 때 한글이 깨어지는 현상이 발생할 수 있다.

### 페이지 개요작성
```html
<meta name="description" content="This Web Site is TIL"> <!-- 페이지의 개요를 작성 -->
```
내 경우에는 페이지의 개요는 지금까지 html문서를 보면서 한번도 본적이 없었는데 신기해서 명시해두었다. 브라우저에서는 어디에서도 표시되지 않지만 검색 사이트 검색결과에서 이 description이 표시가 된다고한다.

### CSS, JAVASCRIPT 불러오기 1
```html
<head> 
<link rel="stylesheet" href="css/style.css"> <!-- css파일을 불러온다 -->
</head>
<body>
<div>
    <p>Hello World</p>
</div>
<script src="js/script.js"></script> <!-- javascript 파일을 불러온다 -->
</body>
```

link 태그를 이용해서 css 파일을 html에 적용시킬수 있고 script 태그를 이용해서 javascript 파일을 적용할 수 있다. 책을 읽으면서 정말 그럴까? 라는 부분이 하나 있었는데 내가 웹 페이지를 개발했던 때에는 css나 javascript의 load되는 순서는 그닥 중요하게 생각하지 않았다. 그런데 위치에 따라 페이지가 load되는 속도에 차이가 있을 수 있다고 책에서 명시가 되어있다. 조금더 자세히 설명을 해보자면

우선 일반적으로 javscript파일의 load가 css보다는 느리므로 css파일을 먼저 명시하는게 좋다. 극단적인 예로 js파일이 엄청나게 커서 5초정도의 시간이 걸리고 그 파일을 불러오는 script태그가 css를 불러오는 link태그보다 위에 있다고 생각해보자

### CSS, JAVASCRIPT 불러오기 2
```html
<head> 
<script src="js/script.js"></script> <!-- javascript 파일을 불러온다 -->
<link rel="stylesheet" href="css/style.css"> <!-- css파일을 불러온다 -->
</head>
<body>
<div>
    <p>Hello World</p>
</div>
</body>
```

만약 위 처럼 될 경우 css가 적용되지 않은채로 페이지가 느리게 load되지 않을까? javascript를 `</body>` 앞에 두게 된다면 일부기능은 제대로 동작되지 않을 수 있으나 (javascript가 제대로 load 되기까지) 페이지가 우선은 빠르게 load 될 것 같다. 하지만 이게 정답이라고는 할 수 없는게 javascript의 구현하고자 하는 내용에 따라 `<head>~</head>` 사이에 기술이 되지 않으면 제대로 동작되지 않는 경우도 있으니 우선 기본적으로는 `</body>`태그 뒤에 기술하는 것을 우선으로하고 만약 작동하지 않을 경우 `<head>`태그 안에 넣는 것을 추천한다. 
