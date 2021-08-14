# 问题收集

1. map和slice的一点区别？  
map声明后需要用make初始化才能使用，否则报异常，而slice不需要，声明后，即可使用，如用append增加成员 。
   
2. int类型会自动化为0，如何表示没有值（类似java中的null)呢？
方法一：用指针*int，可以判断是否为nil，如果不为nil，再取指针变量的值；
方法二：可以用sql.NullInt32,判断Valid为true，再取Int32的值。
方法三：用int表示状态时，可以在设计的时候规定不使用0值，即从1开始表示具体的状态值。避免0和nil的歧义。
   
3. 微信支付如何实现？  
参考库：https://github.com/objcoding/wxpay  
https://github.com/smartwalle/wxpay  
   https://github.com/wleven/wxpay  
   https://github.com/wleven/wxpay  
   
4. 支付宝支付如何实现？  
参考库：https://github.com/smartwalle/alipay  
https://github.com/go-pay/gopay  
https://github.com/ascoders/alipay  
https://github.com/milkbobo/gopay  
   