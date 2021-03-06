package com.example.htttppost;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.app.Activity;
import android.content.Intent;
import android.view.Menu;
import android.view.View;
import android.view.View.OnClickListener;
import android.webkit.WebView;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

public class MainActivity extends Activity implements OnClickListener{

	private Button _sendButton;
	private static final String MAP_URL = "file:///android_asset/map.html";
	private WebView webView;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.main, menu);
        _sendButton = (Button) findViewById(R.id.sendButton);
        assert(_sendButton != null);
        _sendButton.setOnClickListener(this);
        
        //create posChecker Services
        Intent intent = new Intent(MainActivity.this, posChecker.class);
        startService(intent);
        return true;
    }


	@Override
	public void onClick(View arg0) {
        webView = (WebView) findViewById(R.id.webView);
        webView.getSettings().setJavaScriptEnabled(true);       
        webView.loadUrl(MAP_URL);
		
		setContentView(R.layout.test_layout);/*
		final httpService _hs = new httpService();
		EditText _urlText = (EditText) findViewById(R.id.urlText);
		EditText _valueText = (EditText) findViewById(R.id.valueText);
		final String _url = _urlText.getText().toString();
		final String _value = _valueText.getText().toString();
		
		// Handler for httpservice thread
		final Handler mHandler = new Handler() { 

		     public void handleMessage(Message msg) { 
		 		TextView _result = (TextView) findViewById(R.id.resultTextView);
				_result.setText((CharSequence) msg.getData().get("1"));
		     } 
		 };
		 
		 // http thread cant run on main thread
		Thread thread = new Thread(){
		public void run(){
			Bundle bundle = new Bundle();
			bundle.putString("1", _hs.httpServiceGet(_url, _value));	
			Message msg = new Message();
			msg.setData(bundle);
			mHandler.sendMessage(msg);
		}
		};
		thread.start();*/
	}
}
