unit auto_deploy_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, ComCtrls, StdCtrls, Buttons;

type

  { TAutoDeploy }

  TAutoDeploy = class(TFrame)
    ContextPageControl: TPageControl;
    ContextPanel: TPanel;
    DeployContext: TTabSheet;
    FileContextList: TListView;
    GlobalContextList: TListView;
    GlobalSheet: TTabSheet;
    ImageList: TImageList;
    LeftPanel: TPanel;
    BottomPanel: TPanel;
    MainPanel: TPanel;
    NewButton: TSpeedButton;
    OpenDeployButton: TSpeedButton;
    Panel: TPanel;
    SaveButton: TSpeedButton;
    SaveToButton: TSpeedButton;
    AddFileButton: TSpeedButton;
    RemoveButton: TSpeedButton;
    NewFolderButton: TSpeedButton;
    EditButton: TSpeedButton;
    Splitter1: TSplitter;
    Splitter3: TSplitter;
    TopPanel: TPanel;
    TreeView1: TTreeView;
    procedure LeftPanelClick(Sender: TObject);
    procedure MainPanelClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TAutoDeploy }

procedure TAutoDeploy.MainPanelClick(Sender: TObject);
begin

end;

procedure TAutoDeploy.LeftPanelClick(Sender: TObject);
begin

end;

end.

