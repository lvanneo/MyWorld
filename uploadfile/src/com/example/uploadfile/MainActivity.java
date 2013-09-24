package com.example.uploadfile;

import java.io.DataOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.URL;



import com.example.util.Test;

import android.os.Build;
import android.os.Bundle;
import android.os.Environment;
import android.os.StrictMode;
import android.annotation.SuppressLint;
import android.annotation.TargetApi;
import android.app.Activity;
import android.app.AlertDialog;
import android.content.DialogInterface;
import android.view.Menu;
import android.view.SurfaceView;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.TextView;

@TargetApi(Build.VERSION_CODES.GINGERBREAD)
public class MainActivity extends Activity {
	
	private String newName = "psbe.jpg";
    private String uploadFile = "/sdcard/123.zip";
    private String actionUrl = "http://192.168.1.3:8086/upload";
    private TextView mText1;
    private TextView mText2;
    private Button mButton;
    private Button mBtn_start;
    private Button mBtn_end;
    private Button mBtn_fenFile;
    private MovieRecorder movie;
    private SurfaceView surfaceView;
	

	@TargetApi(Build.VERSION_CODES.GINGERBREAD)
	@SuppressLint("NewApi")
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_main);
		
        //-------------------------------------
        //在3.0 及以上版本socket通信不能在UI线程中实现，只能另启线程实现，通过以下连个方法可解决次问题
		StrictMode.setThreadPolicy(new StrictMode.ThreadPolicy.Builder().detectDiskReads().detectDiskWrites().detectNetwork().penaltyLog().build());
		StrictMode.setVmPolicy(new StrictMode.VmPolicy.Builder().detectLeakedSqlLiteObjects().penaltyLog().penaltyDeath().build());
        //-------------------------------------

		this.movie = new MovieRecorder();
		
		this.surfaceView = (SurfaceView) findViewById(R.id.surfaceview);
		
		 /* 设置mButton的onClick事件处理 */    
        mButton = (Button) findViewById(R.id.butUpload);
        mButton.setOnClickListener(new View.OnClickListener()
        {
          public void onClick(View v)
          {
        	  getSDPath();
        	  uploadFile(actionUrl, uploadFile, newName);
          }
        });
        
        this.mBtn_start = (Button) findViewById(R.id.btn_start);
        this.mBtn_start.setOnClickListener(new OnClickListener() {
			
			@Override
			public void onClick(View v) {
				movie.startRecording(surfaceView);
			}
		});
        
        this.mBtn_end = (Button) findViewById(R.id.btn_end);
        this.mBtn_end.setOnClickListener(new OnClickListener() {
			
			@Override
			public void onClick(View v) {
				movie.stopRecording();
			}
		});
        
        this.mBtn_fenFile = (Button)findViewById(R.id.butfenfile);
        this.mBtn_fenFile.setOnClickListener(new OnClickListener() {
			
			@Override
			public void onClick(View v) {
				Test.split(uploadFile, 1024 * 5);
			}
		});
		
		
	}

	@Override
	public boolean onCreateOptionsMenu(Menu menu) {
		// Inflate the menu; this adds items to the action bar if it is present.
		getMenuInflater().inflate(R.menu.main, menu);
		return true;
	}
	
	public void getSDPath() {
		File sdDir = null;
		File sdDir1 = null;
		File sdDir2 = null;
		boolean sdCardExist = Environment.getExternalStorageState().equals(
				android.os.Environment.MEDIA_MOUNTED); // 判断sd卡是否存在
		if (sdCardExist) {
			sdDir = Environment.getExternalStorageDirectory();// 获取跟目录
			sdDir1 = Environment.getDataDirectory();
			sdDir2 = Environment.getRootDirectory();
		}
		System.out.println("getExternalStorageDirectory(): " + sdDir.toString());
		System.out.println("getDataDirectory(): " + sdDir1.toString());
		System.out.println("getRootDirectory(): " + sdDir2.toString());
	}

	
	 /* 上传文件至Server的方法 */
    private void uploadFile(String actionUrl, String uploadFile, String fileName)
    {
      String end = "\r\n";
      String twoHyphens = "--";
      String boundary = "*****";
      try
      {
        URL url =new URL(actionUrl);
        HttpURLConnection con=(HttpURLConnection)url.openConnection();
        /* 允许Input、Output，不使用Cache */
        con.setDoInput(true);
        con.setDoOutput(true);
        con.setUseCaches(false);
        /* 设置传送的method=POST */
        con.setRequestMethod("POST");
        /* setRequestProperty */
        con.setRequestProperty("Connection", "Keep-Alive");
        con.setRequestProperty("Charset", "UTF-8");
        con.setRequestProperty("Content-Type",
                           "multipart/form-data;boundary="+boundary);
        /* 设置DataOutputStream */
        DataOutputStream ds = 
          new DataOutputStream(con.getOutputStream());
        ds.writeBytes(twoHyphens + boundary + end);
        ds.writeBytes("Content-Disposition: form-data; " +
                      "name=\"userfile\";filename=\"" +
                      fileName +"\"" + end);
        ds.writeBytes(end);   

        /* 取得文件的FileInputStream */
        FileInputStream fStream = new FileInputStream(uploadFile);
        /* 设置每次写入1024bytes */
        int bufferSize = 1024;
        byte[] buffer = new byte[bufferSize];

        int length = -1;
        /* 从文件读取数据至缓冲区 */
        while((length = fStream.read(buffer)) != -1)
        {
          /* 将资料写入DataOutputStream中 */
          ds.write(buffer, 0, length);
        }
        ds.writeBytes(end);
        ds.writeBytes(twoHyphens + boundary + twoHyphens + end);

        /* close streams */
        fStream.close();
        ds.flush();

        /* 取得Response内容 */
        InputStream is = con.getInputStream();
        int ch;
        StringBuffer b =new StringBuffer();
        while( ( ch = is.read() ) != -1 )
        {
          b.append( (char)ch );
        }
        /* 将Response显示于Dialog */
        showDialog("上传成功"+b.toString().trim());
        /* 关闭DataOutputStream */
        ds.close();
      }
      catch(Exception e)
      {
        showDialog("上传失败"+e);
      }
    }
    
    /* 显示Dialog的method */
    private void showDialog(String mess)
    {
      new AlertDialog.Builder(MainActivity.this).setTitle("Message")
       .setMessage(mess)
       .setNegativeButton("确定",new DialogInterface.OnClickListener()
       {
         public void onClick(DialogInterface dialog, int which)
         {          
         }
       })
       .show();
    }
	
	
}
