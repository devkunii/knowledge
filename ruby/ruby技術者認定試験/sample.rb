module M
 CONST = "Hello, world"
end
class M::C
 def awesome_method
   CONST
 end
end
p M::C.new.awesome_method
