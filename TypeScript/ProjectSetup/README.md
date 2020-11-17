# Project Setup
The following is the base for TypeScript projects 


## Files 
* [index.html](index.html) - Application HTML file 
* [app.ts](app.ts) - TypeScript application 
* [app.js](app.js) - JavaScript derived from TypeScrpt. 
* [package.json](package.json) - JSON with info regarding the pacakge 

## Steps
1. Write TypeScript code 
2. Convert code to JavaScript 
```
cd ${PROJECT_DIR} 
tsc ${project_file}.ts
```
3. Within the ${PROJECT_DIR}, run `npm init` (NodeJS) in ordeer to **skip** reloading page. 
    * `npm` is part of [Node.js](https://nodejs.org/en/)
    * use default values
4. Run `npm install` install packagess
    * creates node_modules 
    * `--save-dev` option allows for development only
    * [lite-server](https://www.npmjs.com/package/lite-server) - Webapp development server. 
```
npm install --save-dev lite-server 
```
5. Update [package.json](package.json) with line `"start": "lite-server"` bellow "test" 
6. Run `npm start` to run start script 
7. By default, code runs in http://localhost:3000
8. You can then update codee in [app.ts](app.ts) 