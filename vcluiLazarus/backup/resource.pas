unit resource;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls;

type

  { TResource }

  TResource = class(TDataModule)
    ImageList: TImageList;
  private

  public

  end;

var
  Resource: TResource;

implementation

{$R *.lfm}

end.

