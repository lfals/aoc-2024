let compute_distances file =
  Printf.printf "teste";
  let distances = ref 0 in
  let left = ref [||] in
  let right = ref [||] in
  let chan = open_in file in
  try
    while true; do
      let line = input_line chan in
      let parts = String.split_on_char ' ' line in
      let head = List.hd parts in
      let tail = List.nth parts (List.length parts - 1) in
      left := Array.append !left [|head|];
      right := Array.append !right [|tail|];
    done;
  with End_of_file ->
    close_in chan;
    Array.sort String.compare !left;
    Array.sort String.compare !right;
    for i = 0 to (Array.length !left - 1) do
      let left_num = int_of_string (!left).(i) in
      let right_num = int_of_string (!right).(i) in
      let diff = abs (left_num - right_num) in
      distances := !distances + diff
    done;
    !distances

