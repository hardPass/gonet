### HUB 服务器

HUB 作为一个独立的服务运行，各个逻辑服务器(GS)启动时会连接到HUB.      
HUB 管理玩家基础、重要的状态信息的存取变更，如：    

* 排名    
* 保护
* 联盟消息管理   

HUB也承担玩家之间消息转发（不在同一个服务器登陆的情况）。     
HUB只处理来自GS的两类消息：   

1. 来自Game Server 的Call请求(request & ack)     
2. Game Server 间的消息Forward      
