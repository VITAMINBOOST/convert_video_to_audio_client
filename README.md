## 빌드 방법
### For Linux
```
    make linux
```
### For Window
```
    make window
```
---
## 실행 방법
### For Linux
* 다운받은 파일을 바탕화면에 압출을 풀어 놓는다.
* 폴더 안의 xi_mr_trot를 바탕화면에 둔다.
* 터미널을 연다.
* 아래의 것을 붙여 넣는다.
* "<URL 주소 복붙>" 부분에 유튜브 주소를 넣는다. 
```
   cd /Users/`whoami`/Desktop && ./xi_mr_trot -url=<URL 주소 복붙>
```
### For Window
* 다운받은 파일을 아무데나 원하는 곳에 압축을 풀어 놓는다.
* 폴더 안의 xi_mr_trot.exe를 바탕화면에 둔다.
* 같은 위치에서 CMD를 연다.
* 아래의 것을 붙여 넣는다.
* "<URL 주소 복붙>" 부분에 유튜브 주소를 넣는다. 
```
   xi_mr_trot.exe -url=<URL 주소 복붙>
```
---
## ZUM 동영상 URL 추출 방법
### 현재 다운로드 후 변환 가능합니다.
#### 브라우저 주소줄에 있는 URL은 사용 불가능 합니다.
1. ZUM 동영상 재생 페이지에서 개발자 도구 열기
2. 상단의 'Elements'를 클릭 한다.
3. (Window)Ctrl + f 또는 (MAC)Command + f 를 통해 video/mp4를 검색한다.
4. 아래와 같이 검색이 됨.
``` 
<video id="player" playsinline="" tabindex="-1" preload="metadata" 
webkit-playsinline="true" disablepictureinpicture="true" 
controlslist="nodownload" type="video/mp4" title="" 
src="https://zum-channela-c.smartmediarep.com/smc/zum/multi/eng/CA3_000000992093/2f6368616e6e656c612f6e766f642f452f57504744303030303031312f323031392f30322f433139303232313039323733385f323030305f7433332e6d7034/0-0-0/content.mp4?solexpire=1587394894&amp;solpathlen=168&amp;soltoken=6d869c341569da171fab0d9562df8c82&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;soluuid=2213069d-6fb9-47f3-a824-5843452c4827&amp;itemtypeid=33" 
style="visibility: visible;" autoplay="">
                <!-- 플레이어2.0 영역 -->
            </video> 
```
5. 여기서 src 뒤에 있는 주소가 필요합니다.
```
https://zum-channela-c.smartmediarep.com/smc/zum/multi/eng/CA3_000000992093/2f6368616e6e656c612f6e766f642f452f57504744303030303031312f323031392f30322f433139303232313039323733385f323030305f7433332e6d7034/0-0-0/content.mp4?solexpire=1587394894&amp;solpathlen=168&amp;soltoken=6d869c341569da171fab0d9562df8c82&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;soluuid=2213069d-6fb9-47f3-a824-5843452c4827&amp;itemtypeid=33
```
---
## VODA 
### 현재 다운로드 후 변환 가능합니다.
#### 브라우저 주소줄에 있는 URL은 사용 불가능 합니다.
1. Voda 동영상 재생 페이지에서 개발자 도구 열기
2. 상단의 'Elementes'를 클릭 한다.
3. (Window)Ctrl + f 또는 (MAC)Command + f 를 통해  jp_video_0 를 검색한다.
4. 아래와 같이 검색이 됨.
```
<video playsinline="" id="jp_video_0" preload="none" 
src="https://voda-channela-c.smartmediarep.com/smc/voda/multi/eng/CA3_000000989114/2f6368616e6e656c612f6e766f642f452f57504744303030303031312f323031392f30322f433139303231383039303531345f323030305f7433342e6d7034/0-0-0/content.mp4?solexpire=1587402777&amp;soluuid=5e9d2f59d54cf&amp;soltoken=54783050fd7bf1e7138795a959104f7d&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;itemtypeid=34" 
title="[트로트통신 - 임영웅 인터뷰 에피소드 #1] - 삼삼한 날의 콘서트" style="width: 0px; height: 0px;"></video>
```
5. 여기서 src 뒤에 있는 주소가 필요합니다.
```
https://voda-channela-c.smartmediarep.com/smc/voda/multi/eng/CA3_000000989114/2f6368616e6e656c612f6e766f642f452f57504744303030303031312f323031392f30322f433139303231383039303531345f323030305f7433342e6d7034/0-0-0/content.mp4?solexpire=1587402777&amp;soluuid=5e9d2f59d54cf&amp;soltoken=54783050fd7bf1e7138795a959104f7d&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;itemtypeid=34
```
---
## Naver
### 현재 다운로드 후 변환 가능합니다.
#### 브라우저 주소줄에 있는 URL은 사용 불가능 합니다.
1. Naver tv 동영상 재생 페이지에서 개발자 도구 열기(F12)
2. 광고가 지나간 후
3. 상단의 'Elementes'를 클릭 한다.
4. (Window)Ctrl + f 또는 (MAC)Command + f 를 통해 controlslist 를 검색한다.
5. 아래와 같이 검색이 됨.
```
<video aria-hidden="true" id="rmcPlayer_15875460434702209.video" style="left:0;top:0;right:0;bottom:0;width:100%;height:100%;position:absolute;" x-webkit-airplay="" webkit-playsinline="true" playsinline="" disablepictureinpicture="true" controlslist="nodownload" 
data-src="https://naver-mbc-h.smartmediarep.com/smc/naver/adaptive/eng/M01_TZ202004090002/2f6d62632f4174746163682f4d42432f564944454f2f543730303036472f434c49502f636c69705f545a3230323030343039303030325f32303230303430393030303635322e736d696c/0-0-0/playlist.m3u8?solexpire=1587589244&amp;solpathlen=197&amp;soltoken=5d1fc1801c58e49c9d43f2367a57f31e&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;soluuid=09306eb8-33e3-4ab2-9e63-79047e330977" type="video/mp4" src="blob:https://tv.naver.com/53c21878-5a4f-46a6-b568-33b39d4c4596" class="_hide_controls">
<source src="https://naver-mbc-h.smartmediarep.com/smc/naver/adaptive/eng/M01_TZ202004090002/2f6d62632f4174746163682f4d42432f564944454f2f543730303036472f434c49502f636c69705f545a3230323030343039303030325f32303230303430393030303635322e736d696c/0-0-0/playlist.m3u8?solexpire=1587589244&amp;solpathlen=197&amp;soltoken=5d1fc1801c58e49c9d43f2367a57f31e&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;soluuid=09306eb8-33e3-4ab2-9e63-79047e330977" type="video/mp4"></video>
```
6. 여기서 data-src 부분이 필요합니다.
```
data-src="https://naver-mbc-h.smartmediarep.com/smc/naver/adaptive/eng/M01_TZ202004090002/2f6d62632f4174746163682f4d42432f564944454f2f543730303036472f434c49502f636c69705f545a3230323030343039303030325f32303230303430393030303635322e736d696c/0-0-0/playlist.m3u8?solexpire=1587589244&amp;solpathlen=197&amp;soltoken=5d1fc1801c58e49c9d43f2367a57f31e&amp;soltokenrule=c29sZXhwaXJlfHNvbHBhdGhsZW58c29sdXVpZA==&amp;soluriver=2&amp;soluuid=09306eb8-33e3-4ab2-9e63-79047e330977"
```
---
