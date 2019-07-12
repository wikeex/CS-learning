来源：[Vue.js——60分钟快速入门](https://www.cnblogs.com/keepfool/p/5619070.html)，[Vue官方教程](https://cn.vuejs.org/v2/guide/list.html)

***

## v指令

### v-model

双向数据绑定

```vue
<div id="app">
	<p>{{ message }}</p>
	<input type="text" v-model="message"/>
</div>
```

p标签和input表单中任一数据变化都会引起另一元素数据变化。

### v-if

`v-if`是条件渲染指令，它根据表达式的真假来删除和插入元素，它的基本语法如下：

```vue
v-if="expression"
```

expression是一个返回bool值的表达式，表达式可以是一个bool属性，也可以是一个返回bool的运算式。expression为true则渲染该元素，否则不渲染。

### v-show

作用跟v-if类似，但是为false并非不渲染元素，而是设置style="display: none;"来显示隐藏元素。

### v-else

可以用`v-else`指令为`v-if`或`v-show`添加一个“else块”。`v-else`元素必须立即跟在`v-if`或`v-show`元素的后面——否则它不能被识别。

```vue
<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
		<div id="app">
			<h1 v-if="age >= 25">Age: {{ age }}</h1>
			<h1 v-else>Name: {{ name }}</h1>
			<h1>---------------------分割线---------------------</h1>
			<h1 v-show="name.indexOf('keep') >= 0">Name: {{ name }}</h1>
			<h1 v-else>Sex: {{ sex }}</h1>
		</div>
	</body>
	<script src="js/vue.js"></script>
	<script>
		var vm = new Vue({
			el: '#app',
			data: {
				age: 28,
				name: 'keepfool',
				sex: 'Male'
			}
		})
	</script>
</html>
```

### v-for

`v-for`指令基于一个数组渲染一个列表，它和JavaScript的遍历语法相似：

```vue
v-for="item in items"
```

items是一个数组，item是当前被遍历的数组元素。

```vue
<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<title></title>
		<link rel="stylesheet" href="styles/demo.css" />
	</head>

	<body>
		<div id="app">
			<table>
				<thead>
					<tr>
						<th>Name</th>
						<th>Age</th>
						<th>Sex</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="person in people">
						<td>{{ person.name  }}</td>
						<td>{{ person.age  }}</td>
						<td>{{ person.sex  }}</td>
					</tr>
				</tbody>
			</table>
		</div>
	</body>
	<script src="js/vue.js"></script>
	<script>
		var vm = new Vue({
			el: '#app',
			data: {
				people: [{
					name: 'Jack',
					age: 30,
					sex: 'Male'
				}, {
					name: 'Bill',
					age: 26,
					sex: 'Male'
				}, {
					name: 'Tracy',
					age: 22,
					sex: 'Female'
				}, {
					name: 'Chris',
					age: 36,
					sex: 'Male'
				}]
			}
		})
	</script>

</html>
```

### v-bind

`v-bind`指令可以在其名称后面带一个参数，中间放一个冒号隔开，这个参数通常是HTML元素的特性（attribute），例如：`v-bind:class`

```
v-bind:argument="expression"
```

下面这段代码构建了一个简单的分页条，v-bind指令作用于元素的class特性上。
这个指令包含一个表达式，表达式的含义是：高亮当前页。

```vue
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title></title>
		<link rel="stylesheet" href="styles/demo.css" />
	</head>
	<body>
		<div id="app">
			<ul class="pagination">
				<li v-for="n in pageCount">
					<a href="javascripit:void(0)" v-bind:class="activeNumber === n + 1 ? 'active' : ''">{{ n + 1 }}</a>
				</li>
			</ul>
		</div>
	</body>
	<script src="js/vue.js"></script>
	<script>
		var vm = new Vue({
			el: '#app',
			data: {
				activeNumber: 1,
				pageCount: 10
			}
		})
	</script>
</html>
```

vue2中使用略微有差异。

### v-on

`v-on`指令用于给监听DOM事件，它的用语法和v-bind是类似的，例如监听<a>元素的点击事件：

```
<a v-on:click="doSomething">
```

有两种形式调用方法：**绑定一个方法（让事件指向方法的引用），或者使用内联语句。**Greet按钮将它的单击事件直接绑定到greet()方法，而Hi按钮则是调用say()方法。

```vue
<p>
    <!--click事件直接绑定一个方法-->
    <button v-on:click="greet">Greet</button>
</p>
<p>
    <!--click事件使用内联语句-->
    <button v-on:click="say('Hi')">Hi</button>
</p>
```

### v-bind和v-on的缩写

Vue.js为最常用的两个指令`v-bind`和`v-on`提供了缩写方式。**v-bind指令可以缩写为一个冒号，v-on指令可以缩写为@符号。**

```vue
<!--完整语法-->
<a href="javascripit:void(0)" v-bind:class="activeNumber === n + 1 ? 'active' : ''">{{ n + 1 }}</a>
<!--缩写语法-->
<a href="javascripit:void(0)" :class="activeNumber=== n + 1 ? 'active' : ''">{{ n + 1 }}</a>

<!--完整语法-->
<button v-on:click="greet">Greet</button>
<!--缩写语法-->
<button @click="greet">Greet</button>
```

