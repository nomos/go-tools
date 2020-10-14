unit form_list_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, ComCtrls;

type

  { TFormListFrame }

  TFormListFrame = class(TFrame)
    BottomPanel: TPanel;
    ListView1: TListView;
    MainPanel: TPanel;
    Splitter1: TSplitter;
    TopPanel: TPanel;
  private

  public

  end;

implementation

{$R *.lfm}

end.

