```
#### JAVA  

import java.io.*;
import java.net.Socket;
import java.net.UnknownHostException;

public class Test {
    public static void main(String[] args) {
        try {
            //创建一个客户端socket
            Socket socket = new Socket("localhost",1024);
            //向服务器端传递信息
            DataOutputStream ots = new DataOutputStream(socket.getOutputStream());
            String content="用户名：admin;密码：123";

            //java默认是通信进行大端传输，所以这里不用管大小端
            byte[] wb=content.getBytes("utf-8");
            int ln=wb.length;
            byte[] lengthbytes = integerToBytes(ln, 4);

            ots.write(lengthbytes);
            ots.write(wb);
            ots.flush();
            ots.close();
            socket.close();
        } catch (UnknownHostException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }


    public static byte[] integerToBytes(int integer, int len) {
        ByteArrayOutputStream bo = new ByteArrayOutputStream();
        for (int i = 0; i < len; i ++) {
            bo.write(integer);
            integer = integer >> 8;
        }
        return bo.toByteArray();
    }

}


```