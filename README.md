# mypay-export
## Export your pay data
----
>There is a desire to take your pay information and perform your own types of calculations. This tool will allow you to export you pay information into Excel. This utilizes the existing mobile PayHistory API to exctract the data. Users will need to log into the mobile version of the desktop site to utilize the tool.

## Login
Open the browser of your choice and navigate to UltiPro
![Login](/img/login.png)

Take note of your URL and add the value to the config.json
```
{
     "LoginToken": "{{token}}",
     "Url": "https://my.ultimatesoftware.com",
     "NumberOfPays": "10"
}
```

## Mobile
Navigate to the mobile site
![Mobile](/path/to/mobile.png)

## Login Token
Using the developer tools in your browser, look at the cookie for the site. There will be a value for loginToken
![loginToken](/path/to/loginToken.png)
Paste the value into the config.json
```
{
     "LoginToken": "00000000-0000-0000-0000-000000000000",
     "Url": "https://my.ultimatesoftware.com",
     "NumberOfPays": "10"
}
```

## Build and Run
```
	go build
	./mypay-export.exe
```

## Config
Number of pays is defaulted to 10. Max number is 100.