{
  description = "A portable go dev flake!";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { nixpkgs, ... }@inputs: 
	let
    system = "x86_64-linux";
		pkgs = nixpkgs.legacyPackages.${system};
  in 
	{
    devShells.${system} = {
		  default = pkgs.mkShell {
			  nativeBuildInputs = with pkgs; [
					delve
				  	go
					gopls
					templ
				];
				
				hardeningDisable = [ "fortify" ];

				shellHook = "echo Welcome to dev environment";
			};
		};
	};
}
