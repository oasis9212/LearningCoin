# LearningCoin

블록체인 

해쉬 일종의 데이터를 암호화로 치환 시키는 것이라고 생각하면 편하다.

ex) 가가가가가 = hash(가가가가가)  = "paojsdpoaj9018203lojd" 

더 아나가 한글자만 추가하면

ex) 가가가가가나 = hash(가가가가가나)  = "xkicm102pojjuh" 
 SHA-256 해쉬함수를 주로 사용한다.
블록

B1   b1hash = (data+ "") 첫블록의 시작은 이전 블록이 없으므로 이렇게 시작한다. (해쉬는 항상해준다. )
B2   b2hash = (data+b1hash) 이전블록을 물어서 해쉬함수를 돌린다.
B3   b3hash = (data+b2hash) 

만일에  B2가 변경했다고 가정하에  .. 

B2   b2hash = (data+"otherdata"+b1hash)
B3   b3hash = (data+b2hash)  -> 이것에 대한 해쉬값도 변한다. 

POW의 증명 방식: 일종의 Difficulty 와  Nonce 가 있다는 것부터 시작한다.
퀴즈를 제시하면 : 0이 n개의 값이있는 것을 찾아야한다. 
해쉬 함수는 일방향 함수라 데이터를 입력해놓은 이상 수정이 불가능하다.
그래서 데이터+Nonce 해서 0이 n개의 값이 있는 함수를 만들어 준다.  