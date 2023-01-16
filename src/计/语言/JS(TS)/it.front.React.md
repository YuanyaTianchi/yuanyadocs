+++

title = "React"
description = "React"
tags = ["it","front","React"]

+++

# React



## 生命周期

constructor：初始化组件，定义组件的属性（props）和状态(state)

### Mounting阶段

Mounting阶段叫挂载阶段，伴随着整个虚拟DOM的生成

`componentWillMount` :  在组件即将被挂载到页面的时刻执行

`render` : 页面state或props发生变化时执行。在组件state发生变化时自动执行的函数

`componentDidMount`  : 组件挂载完成时被执行

注意：`componentWillMount`和`componentDidMount`这两个生命周期函数，只在页面刷新时执行一次，而`render`函数是只要有state和props变化就会执行

### Updation阶段

`Updation`阶段,也就是组件发生改变的更新阶段，这是React生命周期中比较复杂的一部分，它有两个基本部分组成，一个是`props`属性改变，一个是`state`状态改变

`shouldComponentUpdate`函数会在组件更新之前，自动被执行。它要求返回一个布尔类型的结果，必须有返回值，这里就直接返回一个`true`了（真实开发中，这个是有大作用的）。如果你返回了`false`，这组件就不会进行更新了。 简单点说，就是返回true，就同意组件更新;返回false,就反对组件更新。

`componentWillUpdate`在组件更新之前，但`shouldComponenUpdate`之后被执行。但是如果`shouldComponentUpdate`返回false，这个函数就不会被执行了

`componentDidUpdate`在组件更新之后执行，它是组件更新的最后一个环节

componentWillReceiveProps：子组件接收到父组件传递过来的参数，父组件render函数重新被执行，这个生命周期就会被执行。如果没接收任何的`props`就不会被执行，子组件接收到父组件传递过来的参数，父组件render函数重新被执行，这个生命周期就会被执行。也就是说这个组件第一次存在于Dom中，函数是不会被执行的;如果已经存在于Dom中，函数才会被执行

### Unmounting阶段

`componentWillUnmount`：它是在组件删除时执行



## 脚手架

脚手架就是为项目搭建基本框架，配置依赖一些必须或常用的包（包管理本质上还是webpark）

### create-react-app

create-react-app 是 facebook 官方提供的 react 脚手架工具。create-react-app功能不多，为项目提供最精简的东西

```shell
npm install -g create-react-app #create-react-app安装
```

#### 项目

创建

```shell
E:
cd it/front-end/FrontEndProjects/React
create-react-app demo #创建脚手架项目
```

启动

```shell
cd demo #进入项目目录
npm start #启动项目，会自动在浏览器打开项目首页
```

访问：项目启动默认监听3000端口，http://127.0.0.1:3000

#### 项目目录

- **node_modules**：存放项目依赖包
- **public** ：公共文件夹，有公用模板和图标等一些东西
  - **favicon.ico**：将在浏览器标签页左边显示的图标
  - **index.html**：首页的模板文件
  - **mainifest.json**：移动端配置文件
- **src**：存储所有源代码文件
  - **index.js**：项目入口文件
  - **index.css**：index 的CSS文件
  - **App.js**：一个组件，react的组件都是在js文件中用jsx编写。或者叫模块，react的模式属于模块化编程
  - **serviceWorker.js**：用于写移动端开发，PWA必须用到这个文件，有这个文件就相当于有了离线浏览的功能
- **gitignore**：配置git上传内容。node_modules默认配置了不上传，一般也不上传它
- **package.json**：`webpack`配置和项目包管理文件，项目中依赖的第三方包（包的版本）和一些常用命令配置都在这个里面，脚手架已经为我们配置了好了，无需改动
- **package-lock.json**：版本控制文件。控制锁定所依赖包的版本，保证包环境固定
- **README.md**



## 请求响应

### Axios

Axios，阿可sei偶斯。是一个ajax请求框架，通过npm安装

```shell
npm install -save axios
```

- `npm install xxx`: 安装项目到项目目录下，不会将模块依赖写入`devDependencies`或`dependencies`。
- `npm install -g xxx`: `-g`的意思是将模块安装到全局，具体安装到磁盘哪个位置，要看 `npm cinfig prefix`的位置
- `npm install -save xxx`：`-save`的意思是将模块安装到项目目录下，并在`package`文件的`dependencies`节点写入依赖。
- `npm install -save-dev xxx`：`-save-dev`的意思是将模块安装到项目目录下，并在`package`文件的`devDependencies`节点写入依赖



### Mock

在前后端分离开发中，前端也需要自己模拟响应数据，通常把自己模拟数据这个过程就叫做`mock`。可以用软件自己本地模拟数据，也可以用一些mock平台，推荐淘宝的rap2，http://rap2.taobao.org/



## JSX

### 程序入口

```jsx
/* index.js
脚手架创建的项目，一般默认以 index.js 作为入口文件
在一个html中，React对其组件标签外的任何其它DOM不产生任何影响，在html还是常规引入js、css都可以，但是一般不建议这么做，为了统一技术栈*/

/* React必要组件 */
import React from 'react';
import ReactDOM from 'react-dom';

/* 自定义组件。引入组件可以省略.js；引入时的组件名定义尽量与文件名一致 */
import Xiaojiejie from './component/Xiaojiejie';

/* css */
import './css/style.css'

/* 通过组件ReactDOM来render我们传入的自定组件，添加到html页面的根dom节点中*/
ReactDOM.render(<Xiaojiejie />, document.getElementById('root'))
```

### 父组件

```jsx
/* 自定义组件：React的组件，就是在JS文件中使用JSX语法编写的
约定：自定义组件须首写字母大写，JSX则须首字母小写
JSX是JavaScript和XML结合的一种格式，由React发明，能便利的用HTML语法来创建虚拟DOM，当遇到`<`时就作HTML解析，遇到`{`时就作JavaScript解析 */

import React, { Component, Fragment } from 'react'
import axios from 'axios'
import Service from './Service'
import Boss from './Boss'

class Xiaojiejie extends Component {
    /* constructor：构造函数 */
    constructor(props) {
        super(props) //调用父类构造函数。固定写法

        /* state：组件状态。其实就是组件数据。组件state决定了组件render出来的内容 */
        this.state = {
            inputValue: '摩擦', //input框的值
            serviceList: [], //服务列表
            addServiceButton: '增加服务', //添加服务按钮内容
        }

        /* 组件绑定：为函数绑定this，this即当前这个组件实例，否则函数内this并非指向该组件，会指向错误 */
        this.inputChange = this.inputChange.bind(this)
        this.addService = this.addService.bind(this)
        this.deleteItem = this.deleteItem.bind(this)
    }

    /* render */
    render() {
        return (
            /* 属性冲突：jsx中有js和xml，为了避免关键字冲突，jsx为冲突关键字定义了替代字
            在jsx中for属性用htmlFor代替，避免与js循环关键字for混淆，否则报warning
            在jsx中class属性用className代替，避免与js类定义关键字class混淆，否则报warning */

            /* 组件render时必须要有且仅有一个最外层标签，否则报错。一般使用<div>
            但是可能有不能有外层dom节点的情况，比如在作Flex布局的时候
            react16提供了Fragment组件：用<Fragment>代替<div>，渲染到浏览器中将不再有该外层dom节点（元素）*/
            <Fragment>
                {/* 注意从进入标签开始，注释要写在 {} 范围内 */}

                <div>
                    {/* label是html中非常有用的一个辅助标签，一般是for="要辅助的dom节点id"，这里用htmlFor */}
                    <label htmlFor="jspang">加入服务：</label>

                    <input id="摩擦" className="input"
                        /* 数据绑定：通过{}绑定，算是js代码的标识，就是在JSX中使用js代码。绑定以后，无法在input框中修改内容，React是单向绑定的，只能通过修改组件state来改变dom节点，所以通过事件来修改组件state即可实现双向绑定（vue是v-model直接双向绑定的） */
                        value={this.state.inputValue}

                        /* 组件绑定：为函数绑定this，this即当前这个组件实例，否则函数内this会指向错误 
                        通过.bind(this)也可以绑定this组件：onChange={this.inputChange.bind(this)}
                        一般建议在constructor中绑定即可，性能会高一些 */
                        onChange={this.inputChange}
                    />

                    {/* 数据显示
                    纯文本：<div>{this.state.xxx}</div>
                    html：<div dangerouslySetInnerHTML={{__html: xxx}}></div> */}
                    <button onClick={this.addService}> {this.state.addServiceButton} </button>
                </div>

                <ul
                    /* 通过ref属性：ref={(ele) => { this.xxx = ele }}，可以用自定义的名字绑定dom元素到this，绑定后可以通过 this.xxx 直接在任何地方调用这个dom元素。但一般不建议使用ref，直接操作dom元素会出现一些问题，因为setState是异步过程的，假如在某个函数中写setState修改组件state中的一个数据，代码紧接着又通过this直接操作了这个数据绑定的dom元素，就可能出现并发修改或类似脏读等的问题。虽然setState可以通过第二个参数设置回调函数，在回调函数中直接操作dom元素不存在并发问题，但是还是不建议使用，因为React的是数据驱动的，由state决定dom元素，这种只关心数据的理念就是React最大的优点，排除了同时操作state和dom的复杂性，直接操作dom元素有违这种理念*/
                    ref={(ul) => { this.ul = ul }}>


                    {this.state.serviceList.map(
                        /* 这里的函数与render函数本质上是相通的，都是返回一个dom节点到父节点下，所以规则也是一样的，比如也必须要有且仅有一个最外层标签 */
                        (item, index) => {
                            return (
                                <Fragment key={index}>
                                    <Service
                                        /* 每个标签都需要唯一key，否则报 warning */
                                        key={'son' + index + item}
                                        /* 数据传递：父组件向子组件传递数据，通常通过属性传递 */
                                        service={item}
                                        index={index}
                                        /* 函数传递是一样的，不过仍然要注意绑定this */
                                        deleteItem={this.deleteItem}
                                    />
                                </Fragment>
                            )
                        })
                    }
                </ul>

                {/* 演示动画的组件 */}
                <Boss serviceList={this.state.serviceList}></Boss>
            </Fragment>
        )
    }

    /* componentDidMount */
    componentDidMount() {
        /* Axios请求：需要 import axios from 'axios'
        建议在生命周期函数componentDidMount函数里执行发送ajax请求
        在render函数里请求，会出现很多问题，比如一直循环渲染，因为render函数会经常在组件更新时调用
        在componentWillMount函数里请求，使用RN时会有冲突 */
        axios.get('http://rap2.taobao.org:38080/app/mock/258607/test').then((result) => {
            console.log('axios 获取数据成功:' + JSON.stringify(result))
            this.setState({
                serviceList: result.data.serviceList
            })
        }).catch((error) => {
            console.log('axios 获取数据失败' + error)
        })
    }

    /* 输入框改变事件 */
    inputChange(e) {
        /* 组件状态修改
        this.state.inputValue = e.target.value; 无法直接修改组件state
        React必须通过this调用setState函数修改组件state，这样React才知道组件state改变了以渲染组件
        setState是异步过程的 */
        this.setState({
            inputValue: e.target.value
        })
    }

    /* 增加服务 */
    addService(e) {
        this.setState(
            {
                /* 扩展运算符...是ES6语法，就是把数组进行了分解，然后再进行组合，形成新的数组 */
                serviceList: [...this.state.serviceList, this.state.inputValue]
            },
            /* setState可以通过第二个参数设置回调函数 */
            () => {
                console.log(this.ul.querySelectorAll('li').length)
            }
        )
    }

    /* 删除服务 */
    deleteItem(index) {
        let serviceList = this.state.serviceList
        serviceList.splice(index, 1)
        this.setState({
            serviceList: serviceList
        })
    }
}

export default Xiaojiejie
```

### 子组件

```jsx
/* Service.js */

import React, { Component } from 'react'; //imrc
import PropTypes from 'prop-types'

class Service extends Component { //cc
    constructor(props) {
        super(props)
        this.handleClick = this.handleClick.bind(this)
    }

    /* 不添加这个方法时，会存在一个性能问题，就是即使只是在输入框中输入一个字符，也会导致render被调用
    简单粗暴 shouldComponentUpdate(){ return false; } 可以解决问题，但是如果想要改变属性值，也会无法render
    所以，当属性改变时，让其return true，其它情况return false */
    shouldComponentUpdate(nextProps, nextState) { //（变化后的属性，变化后的状态）
        if (nextProps.content !== this.props.content) {
            return true
        } else {
            return false
        }
    }

    render() {
        console.log('child-render')
        return (
            <li onClick={this.handleClick}>
                {/* 通过 this.props.属性名 获取父组件传递的属性 */}
                {this.props.avname}-为你做-{this.props.service}
            </li>
        );
    }

    handleClick() {
        /* 在子组件中调用父组件的函数 */
        this.props.deleteItem(this.props.index)
    }
}

export default Service;

/* prop-types：import PropTypes from 'prop-types' */
/* propTypes：用于校验数据类型，如果父组件传递过来的数据不符合指定数据类型，则报warning */
Service.propTypes = {
    content: PropTypes.string,
    deleteItem: PropTypes.func,
    index: PropTypes.number.isRequired, //isRequired表示这是一个必需的属性，父组件必须传递一个属性值，否则warning
}
/* defaultProps：用于为props设置默认值 */
Service.defaultProps = {
    avname: '松岛枫'
}
```

### 动画

```jsx
import React, { Component } from 'react';
import { CSSTransition, TransitionGroup } from 'react-transition-group'

class Boss extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isShow: true,
            isShow2: true
        }
        this.toToggole = this.toToggole.bind(this)
        this.toToggole2 = this.toToggole2.bind(this)
    }
    render() {
        return (
            <div>
                {/* css3 */}
                <div className={this.state.isShow ? 'show2' : 'hide2'}>BOSS级人物-孙悟空</div>
                <div><button onClick={this.toToggole}>召唤Boss</button></div>

                {/* CSSTransition */}
                <CSSTransition
                    in={this.state.isShow2} //用于判断是否出现的状态
                    timeout={2000}          //动画持续时间
                    classNames="boss-text"  //className值，防止重复，所以有一个s
                    unmountOnExit //unmountOnExit使：动画结束时自动删除dom节点，动画开始时自动添加dom节点。这是纯CSS动画做不到的
                >{/* 这里在浏览器F5的时候没有动画效果，点击有动画效果，不知道为啥，暂时不管了 */}
                    <div>BOSS级人物-孙悟空</div>
                </CSSTransition>
                <div><button onClick={this.toToggole2}>召唤Boss2</button></div>

                <ul>
                    {/* TransitionGroup
                    动画组。作为一组CSSTransition动画组件的父节点，实际动画效果仍然是由CSSTransition定义 */}
                    <TransitionGroup>
                        {this.props.serviceList.map((item, index) => {
                            return (
                                /* 这里只能由CSSTransition作为返回的最外层dom节点，个人猜测是CSSTransition必须是TransitionGroup的直接子节点。CSSTransition内有且仅能有一个子节点，不能有任何更多子节点甚至文字 */
                                <CSSTransition
                                    timeout={1000}
                                    classNames='boss-text'
                                    unmountOnExit
                                    appear={true}
                                    key={index + item}
                                >
                                    <li>{item}</li>
                                </CSSTransition>
                            )
                        })}
                    </TransitionGroup>
                </ul>
            </div>
        )
    }

    toToggole() {
        this.setState({
            isShow: this.state.isShow ? false : true
        })
    }

    toToggole2() {
        this.setState({
            isShow2: this.state.isShow2 ? false : true
        })
    }
}

export default Boss;
```

#### css3

纯css实现的动画，无法直接或间接改变dom节点的存在

##### transition

```css
/* transition
是纯css3实现的动画，只能作一些最简单的动画。以淡入显示、淡出隐藏为例 */

.show {
    opacity: 1;
    transition: all 1.5s ease-in;
}

.hide {
    opacity: 0;
    transition: all 1.5s ease-in;
}
```

##### keyframes

```css
/* keyframes
可以实现关键帧动画。此属性与`animation`属性是密切相关的，`keyframes`译成中文就是关键帧，最早接触这个关键帧的概念是字flash中，现在Flash已经退出历史舞台了。他和`transition`比的优势是它可以更加细化的定义动画效果
相比于比如上面transition的显示隐藏动画，keyframes还可以设置透明度、颜色等
也是纯css3实现的动画，只能实现很简单的动画效果，一些复杂的动画最好还是使用别人造好的轮子*/

@keyframes hide-item {
    0% {
        opacity: 1;
        color: green;
    }
    50% {
        opacity: 0.5;
        color: red;
    }
    100% {
        opacity: 0;
        color: yellow;
    }
}

@keyframes show-item {
    0% {
        opacity: 0;
        color: yellow;
    }
    50% {
        opacity: 0.5;
        color: red;
    }
    100% {
        opacity: 1;
        color: green;
    }
}

.hide2 {
    /* animation：使用动画的关键词，后边跟上keyframes动画名称
    forwards属性用来控制停止到动画最后一帧，否则动画执行完会恢复原状 */
    animation: hide-item 2s ease-in forwards; 
}

.show2 {
    animation: show-item 2s ease-in forwards;
}
```



#### 动画组件

动画组件不止css，还包括了js操作组件的逻辑，可以影响dom节点增删

##### react-transition-group

react-transition-group是react官方提供的动画过渡库，有着完善的API文档 https://github.com/reactjs/react-transition-group、https://reactcommunity.org/react-transition-group/ 。通过npm安装

```shell
npm install --save react-transition-group
```

```css
/* react-transition-group
xxx-enter: 进入（入场）前的CSS样式；
xxx-enter-active:进入动画直到完成时之前的CSS样式;
xxx-enter-done:进入完成时的CSS样式;
xxx-exit:退出（出场）前的CSS样式;
xxx-exit-active:退出动画知道完成时之前的的CSS样式。
xxx-exit-done:退出完成时的CSS样式。*/

.boss-text-enter {
    opacity: 0;
}

.boss-text-enter-active {
    opacity: 1;
    transition: opacity 2000ms;
}

.boss-text-enter-done {
    opacity: 1;
}

.boss-text-exit {
    opacity: 1;
}

.boss-text-exit-active {
    opacity: 0;
    transition: opacity 2000ms;
}

.boss-text-exit-done {
    opacity: 0;
}
```



## 其它

### 调试工具

React Developer Tools，火狐、chrome都有。添加插件过后在浏览器F12中就多了react的选项卡，注意在react选项卡右边的设置按钮中勾选highlight update when component render，即在组件render时高亮修改的部分

### 快速代码插件

搜索VSCode的插件：Simple React Snippets，安装过后几个缩写字母即可快速补全常用代码。（不使用插件，自己根据习惯配置vscode的快速提示生成也可行）

imrc：快速生成最常用的 import 代码

```react
import React, { Component } from 'react';
```

cc：快速生成组件定义代码

```react
class  extends Component {
    state = {  }
    render() { 
        return (  );
    }
}
export default ;
```



## Ant Design

https://ant.design

`Ant Design`是一套面向企业级开发的UI框架，视觉和动效作的很好。阿里开源的一套UI框架，它不只支持`React`，还有`ng`和`Vue`的版本。

```shell
npm install antd --save
```

