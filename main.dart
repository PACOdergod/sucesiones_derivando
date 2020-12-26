import 'dart:math' as math;

void main(args) {
  // todas las funciones por comodidad
  // usaran un arreglo de 4 numeros

  // x := [4]int{}
  // List<num> x = new List(4);

  // PRIMER PASO
  // una sucesion cuyo primer termino sea uno
  // con esta formula obtendremos 1 en las 4 posiciones
  // primerPaso(x);

  // SEGUNDO PASO
  // una funcion que regrese [1,7]
  // segundoPaso(x);

  // TERCER PASO
  // una funcion que regrese [1,7,11]
  // tercerPaso(x);

  // CUARTO PASO
  // una funcion que regrese [1,7,11,pi]
  // cuartoPaso(x);

  // funcion que recibe los numeros y contruya la funcion lineal
  /// para obtener la funcion hay que basarce en la posicion del numero
  /// esto sera el index
  /// y su formula sera el conjunto de operaciones que me den el numero que quiero
  /// obtener en esa posicion
  /// como el index en computacion empieza en cero hay que incluirlo para evitar problemas

  List<int> x = [0, 1, 2, 3, 4, 5, 10];
  crearFuncion(x);
}

void crearFuncion(List<int> sucesion) {
  List<int> _listaComprobada = [];
  for (var n = 1; n < sucesion.length; n++) {
    /// ahora si tenomos n
    /// para una funcion lineal hay que obtener la diferancia entre
    /// los valores en cada posicion

    var valorIndex = sucesion[n] - sucesion[n - 1];
    // print(sucesion[n]);
    _listaComprobada.add(valorIndex);
  }
  // print(_listaComprobada);
  comprobarIgualdad(_listaComprobada);
}

void comprobarIgualdad(List<int> lista) {
  List<int> _listaIgualdad = [];

  if (lista.length > 1) {
    for (var i = 1; i < lista.length; i++) {
      if (lista[i] == lista[i - 1]) {
        // print('bien');
        _listaIgualdad.add(lista[i]);
      } else {
        // print('mal');
        _listaIgualdad.add(lista[i]);
      }
    }
    // print(_listaIgualdad);

    comprobarIgualdad(_listaIgualdad);
  } else {
    if (lista.first == 1) {
      print('Sucesion lineal');
    } else {
      print('no lineal');
    }
  }
  // print(_listaIgualdad);
}

//
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
///
void primerPaso(List<num> x) {
  // en el video da la formula an = 1
  // n se refiere a la posicion del numero
  // esto lo podemos usar facil como el index de un arreglo
  // hay que iterar en cada index del arreglo y le asignaremos el valor de 1

  for (var item in x) {
    item = 1;
    print(item);
  }
}

void segundoPaso(List<num> x) {
  // formula : bn =(n-1)*6+1
  // n es la posicion del numero, osea el index+1 de un arreglo
  // b es el valor en cada posicion
  for (var i = 0; i < x.length; i++) {
    x[i] = i * 6 + 1;
  }
  print(x);
}

void tercerPaso(List<num> x) {
  // e el video usa cn = -(n-1)(n-2) + (n-1)*6+1

  for (var i = 0; i < x.length; i++) {
    x[i] = -i * (i - 1) + i * 6 + 1;
  }
  print(x);
}

void cuartoPaso(List<num> x) {
  // funcion: ((3.1416-13)/6)* (n-1)(n-2)(n-3)+(-1)*(n-2)+(n-1)*6+1
  for (var i = 0; i < x.length; i++) {
    x[i] = ((math.pi - 13) / 6) * i * (i - 1) * (i - 2) +
        (-i) * (i - 1) +
        i * 6 +
        1;
  }
  print(x);
}
