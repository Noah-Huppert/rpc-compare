using Go = import "/go.capnp";

@0xe3d94c5bbb366af5;

$Go.package("calc");
$Go.import("calc");

interface Calculator {
	add @0 (numbers: List(Int32)) -> (result: Int32);
}
