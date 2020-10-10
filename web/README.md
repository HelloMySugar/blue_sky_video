# Video on demand based on Vue, ElementUI

**Technology used**

```
UI framework: EelementUI
Network request: axios
Video player component: vue-player
```

**basic skills**

```
1. Video browsing
2. Video search
3. Video Management
4. Video upload
5. Video tags
```

**Back-end interface**

Open vue.config.js to configure the service interface address:

```javascript
devServer: {
		port: 8000, // Port number of the project
		// Cross-domain proxy configuration
		proxy: {
			'/api': {
				target: 'http://localhost:8888', 				ws: true,//是否代理websockets
				changeOrigin: true,   //Set same origin, default false, do you need to change the original host header to the target URL
				pathRewrite: {
						'^/api': ''
				}		
			}
		}
	},
```

### Installation dependencies

```
npm install
```

### Start service

```
npm run serve 
```

After starting the service, open the browser and visit http://localhost:8000

### Project packaging

```
npm run build
```
